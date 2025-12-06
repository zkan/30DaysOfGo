package main

import (
	"net/http"
	"strings"
)

type PlayerStore interface {
	recordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if r.Method == http.MethodPost {
		p.store.recordWin(player)
		w.WriteHeader(http.StatusAccepted)
	}
}
