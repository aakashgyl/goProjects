package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MainUrlBody struct {
	Url string `json:"url"`
}

type UrlRespBody struct {
	ShortenedUrl string `json:"short_url"`
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		shortUrl := r.URL.Path
		shortUrl = shortUrl[1:]
		url, err := s.urlService.GetMainUrl(shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mainUrlRespBody := MainUrlBody{
			Url: url,
		}
		respBodyStr, _ := json.Marshal(mainUrlRespBody)
		fmt.Fprintf(w, "%s", respBodyStr)

	case "POST":
		reqBody := MainUrlBody{}
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		shortenedUrl, err := s.urlService.CreateShortenedUrl(reqBody.Url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		respBody := UrlRespBody{ShortenedUrl: "localhost:" + port + "/" + shortenedUrl}
		respBodyStr, _ := json.Marshal(respBody)
		fmt.Fprintf(w, "%s", respBodyStr)

	default:
		w.Write([]byte(`{"message": "not found"}`))
	}
}
