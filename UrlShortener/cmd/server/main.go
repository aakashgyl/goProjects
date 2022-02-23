package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aakashgyl/UrlShortener/pkg/url_shortener"
)

const port = "8888"

type server struct {
	urlService url_shortener.UrlService
}

func main() {
	urlService, err := url_shortener.CreateUrlServiceObject()
	if err != nil {
		log.Fatalf("Failed to initialize url operator: %s", err.Error())
	}

	s := &server{
		urlService: urlService,
	}

	http.Handle("/", s)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
