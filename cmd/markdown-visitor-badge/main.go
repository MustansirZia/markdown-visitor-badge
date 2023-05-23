package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MustansirZia/markdown-visitor-badge/api"
	"github.com/MustansirZia/markdown-visitor-badge/configProvider"
)

func main() {
	config, err := configProvider.NewConfigProvider().Provide()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", api.Handler)
	fmt.Printf("Server is listening on port %d...", config.Port)
	http.ListenAndServe(":8080", nil)
}
