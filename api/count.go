package api

import (
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/MustansirZia/markdown-visitor-badge/configProvider"
	"github.com/MustansirZia/markdown-visitor-badge/requestParser"
	"github.com/MustansirZia/markdown-visitor-badge/store"
	"github.com/MustansirZia/markdown-visitor-badge/svgRenderer"
)

var counter store.Counter
var once sync.Once
var renderer svgRenderer.Renderer
var parser requestParser.RequestParser

// Handler - Handles the HTTP request.
func Handler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/" && r.URL.Path != "/api/count") || r.Method != http.MethodGet {
		writeNotFoundResponse(w)
		return
	}
	once.Do(func() {
		if config, err := configProvider.NewConfigProvider().Provide(); err != nil {
			writeErrorResponse(w, err)
		} else {
			counter = store.NewCounter(config)
		}
		renderer = svgRenderer.NewRenderer()
		parser = requestParser.NewParser()
	})
	if requestParams, err := parser.Parse(r); err != nil {
		if err == requestParser.ErrInvalidKey {
			writeBadRequestResponse(w, err)
			return
		}
		writeErrorResponse(w, err)
		return
	} else {
		if counter == nil {
			return
		}
		if count, err := counter.IncrementAndGet(r.Context(), requestParams.Key); err != nil {
			writeErrorResponse(w, err)
		} else {
			w.Header().Set("Content-Type", "image/svg+xml;")
			if err := renderer.RenderBadge(
				w,
				svgRenderer.BadgeParams{
					Count:        count,
					CountBgColor: requestParams.CountBgColor,
					CountColor:   requestParams.CountColor,
					Label:        requestParams.Label,
					LabelBgColor: requestParams.LabelBgColor,
					LabelColor:   requestParams.LabelColor,
				},
			); err != nil {
				writeErrorResponse(w, err)
			}
		}
	}
}

func writeBadRequestResponse(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Bad Request</h1><br/>Details: %s", err.Error())
}

func writeErrorResponse(w http.ResponseWriter, err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Internal Server Error</h1><br/>Details: %s", err.Error())
}

func writeNotFoundResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>Not Found</h1><br/>Details: Nothing to see here.")
}
