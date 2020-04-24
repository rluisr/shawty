// Package handlers provides HTTP request handlers.
package handlers

import (
	"log"
	"net/http"

	"github.com/rluisr/shawty/storages"
)

func EncodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.URL.Query().Get("url"); url != "" {
			code, err := storage.Save(url)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("err: " + err.Error()))
				return
			}

			_, err = w.Write([]byte(code))
			if err != nil {
				log.Printf("err: %v\n", err)
				return
			}
		}
	}

	return http.HandlerFunc(handleFunc)
}

func DecodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/dec/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			if err != nil {
				log.Printf("err: %v\n", err)
				return
			}
			return
		}

		_, err = w.Write([]byte(url))
		if err != nil {
			log.Printf("err: %v\n", err)
			return
		}
	}

	return http.HandlerFunc(handleFunc)
}

func RedirectHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/red/"):]

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			_, err = w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			if err != nil {
				log.Printf("err: %v\n", err)
				return
			}
			return
		}

		http.Redirect(w, r, url, 301)
	}

	return http.HandlerFunc(handleFunc)
}
