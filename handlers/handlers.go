// Package handlers provides HTTP request handlers.
package handlers

import (
	"log"
	"net/http"

	"github.com/rluisr/shawty/storages"
)

func EncodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			log.Print("save url: " + url)
			_, err := w.Write([]byte(storage.Save(url)))
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
		log.Print("read saved url for code: " + code)

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
		log.Print("redirect to saved url for code: " + code)

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

		http.Redirect(w, r, string(url), 301)
	}

	return http.HandlerFunc(handleFunc)
}
