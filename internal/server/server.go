package server

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"tjan-src-rank/internal/src"
	"tjan-src-rank/internal/util"
)

const timeout = 5

type Handler struct {
	SrcAPI *src.API
}

func (handler *Handler) rankHandler(w http.ResponseWriter, _ *http.Request) {
	id, rank, time, err := handler.SrcAPI.GetRun()
	if err != nil {
		if errors.Is(err, src.ErrNoRunsFound) {
			_, _ = io.WriteString(w, "Keine Runs gefunden")
		} else {
			_, _ = io.WriteString(w, "@punkystone irgendwas kaputt DinkDonk")
		}
	}
	_, _ = io.WriteString(w, fmt.Sprintf("Tjans aktueller Platz: %d (%s): https://www.speedrun.com/mc/runs/%s", rank, util.FormatSeconds(time), id))
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
