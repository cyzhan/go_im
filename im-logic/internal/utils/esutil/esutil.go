package esutil

import (
	"bytes"
	"context"
	"crypto/tls"
	b64 "encoding/base64"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	esClient         *elasticsearch.Client
	esAddress        = os.Getenv("ELASTIC_ADDRESS")
	customHttpClient *http.Client
	basicAuth        string
)

func InitES() {
	cfg := elasticsearch.Config{
		Addresses: []string{os.Getenv("ELASTIC_ADDRESS")},
		Username:  os.Getenv("ELASTIC_USER"),
		Password:  os.Getenv("ELASTIC_PASSWORD"),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true, //添加这一行跳过验证
			},
		},
	}

	var err error
	if esClient, err = elasticsearch.NewClient(cfg); err != nil {
		log.Fatalf(err.Error())
	}
	d := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	customHttpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return d.DialContext(ctx, network, addr)
			},
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		},
	}

	s := os.Getenv("ELASTIC_USER") + ":" + os.Getenv("ELASTIC_PASSWORD")
	basicAuth = b64.StdEncoding.EncodeToString([]byte(s))

	log.Printf("package init ok: esutil. version = %s", elasticsearch.Version)
}

func QueryRequest(reqBody string, index string) ([]byte, error) {
	path := esAddress + "/" + index + "/_search"
	req, err := http.NewRequest("POST", path, bytes.NewReader([]byte(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth)
	res, err := customHttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	val, err := ioutil.ReadAll(res.Body)
	return val, nil
}
