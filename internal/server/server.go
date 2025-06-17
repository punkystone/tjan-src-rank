package server

import (
	"io"
	"net/http"
	"time"
	"tjan-src-rank/internal/src"

	"github.com/rs/zerolog/log"
)

const timeout = 5

type Handler struct {
	SrcAPI *src.API
}

func (handler *Handler) rankHandler(w http.ResponseWriter, _ *http.Request) {
	rank := handler.SrcAPI.GetRank("TjanTV")
	_, err := io.WriteString(w, rank)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write response")
	}
}

func StartServer(srcAPI *src.API) error {
	handler := &Handler{
		SrcAPI: srcAPI,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.rankHandler)
	server := &http.Server{
		Addr:              ":80",
		Handler:           mux,
		ReadHeaderTimeout: timeout * time.Second,
	}

	return server.ListenAndServe()
}
