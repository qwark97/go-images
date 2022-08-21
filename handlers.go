package main

import (
	"encoding/json"
	"net/http"

	"example.com/app/storage"
)

func NewHandlers(s *storage.Storage) *Handlers {
	return &Handlers{
		s: s,
	}
}

type Handlers struct {
	s *storage.Storage
}

func (h *Handlers) Post(w http.ResponseWriter, r *http.Request) {
	var err error
	var data storage.CreateData
	var resp storage.CreateResp

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errResp(err, w)
		return
	}

	if resp, err = h.s.Create(data); err != nil {
		errResp(err, w)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	var err error
	var resp []storage.ReadResp

	if resp, err = h.s.Read(); err != nil {
		errResp(err, w)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func errResp(err error, w http.ResponseWriter) {
	if err == nil {
		return
	}
	status := http.StatusBadRequest
	w.WriteHeader(status)

	type resp struct {
		Status int    `json:"status"`
		Msg    string `json:"message"`
	}
	json.NewEncoder(w).Encode(resp{Status: status, Msg: err.Error()})
}
