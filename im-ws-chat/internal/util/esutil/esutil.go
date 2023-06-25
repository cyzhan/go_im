package esutil

import (
	"bytes"
	"context"
	"crypto/tls"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"imws/internal/constant/esindex"
	"imws/internal/model/errcode"
	"imws/internal/model/im"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	esClient         *elasticsearch.Client
	esAddress        = os.Getenv("ELASTIC_ADDRESS")
	customHttpClient *http.Client
	basicAuth        string
)

func init() {
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

func panicIfError(err error) {
	if err != nil {
		panic(errcode.InternalError(err))
	}
}

type chatSearchVO struct {
	Hits *chatSearchHits `json:"hits"`
}

type chatSearchHits struct {
	Hits []*chatSearchWrapper `json:"hits"`
}

type chatSearchWrapper struct {
	Source *im.MessageEntity `json:"_source"`
}

type textingSearchVO struct {
	Hits *textingSearchHits `json:"hits"`
}

type textingSearchHits struct {
	Hits []*textingSearchWrapper `json:"hits"`
}

type textingSearchWrapper struct {
	Source *im.TextingMessage `json:"_source"`
}

func SaveOne(docID string, v any, esIndex string) {
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(v)
	response, err := esClient.Create(esIndex, docID, body)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Printf("%v", response)
}

func SaveMany(msgs map[string]*any, esIndex string) {
	if len(msgs) == 0 {
		log.Printf("len(msgs) == 0")
		return
	}
	body := &bytes.Buffer{}
	for id, value := range msgs {
		meta := []byte(fmt.Sprintf(`{"index":{"_id":"%s"}}%s`, id, "\n"))
		data, err := json.Marshal(value)
		if err != nil {
			log.Printf("json.Marshal fail")
			return
		}
		data = append(data, "\n"...)
		body.Grow(len(meta) + len(data))
		body.Write(meta)
		body.Write(data)
	}

	respone, _ := esClient.Bulk(body, esClient.Bulk.WithIndex(esIndex))
	if respone.StatusCode != 200 {
		log.Printf("%v", respone)
	}
}

func queryRequest(reqBody string, index string) ([]byte, error) {
	path := esAddress + "/" + index + "/_search"
	req, err := http.NewRequest("POST", path, bytes.NewReader([]byte(reqBody)))
	// req, err := http.NewRequest("POST", esAddress+"/jim-message/_search", bytes.NewReader([]byte(reqBody)))

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

func GetChatMsg(chatID int64, pointer int64) []*im.MessageEntity {
	var resBytes []byte
	var err error
	if pointer == 0 {
		resBytes, err = queryRequest(script.LatestChatMsg(chatID), esindex.CHAT)
	} else {
		resBytes, err = queryRequest(script.OldChatMsg(chatID, pointer), esindex.CHAT)
	}
	panicIfError(err)

	var k []*chatSearchWrapper
	obj := &chatSearchVO{
		Hits: &chatSearchHits{
			Hits: k,
		},
	}
	err = json.Unmarshal(resBytes, obj)
	panicIfError(err)
	count := len(obj.Hits.Hits)
	values := make([]*im.MessageEntity, count)
	for i, item := range obj.Hits.Hits {
		values[i] = item.Source
	}
	return values
}

func GetTextingMsg(msgId []int64) []*im.TextingMessage {
	var resBytes []byte
	var err error
	resBytes, err = queryRequest(script.fetchTextingMsg(msgId), esindex.TEXTING)
	panicIfError(err)

	var k []*textingSearchWrapper
	obj := &textingSearchVO{
		Hits: &textingSearchHits{
			Hits: k,
		},
	}
	err = json.Unmarshal(resBytes, obj)
	panicIfError(err)
	count := len(obj.Hits.Hits)
	values := make([]*im.TextingMessage, count)
	for i, item := range obj.Hits.Hits {
		values[i] = item.Source
	}
	return values
}
