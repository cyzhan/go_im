package service

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	System = &systemHandler{"v0.1.0"}
)

type systemHandler struct {
	version string
}

func (hd *systemHandler) GetVersion(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"version": hd.version,
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("get version err: %s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
