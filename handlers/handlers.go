// Package handlers provides HTTP request handlers.
package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/didip/shawty/storages"
)

type PageData struct {
	BaseUrl   string `"http://localhost:8080"`
	Shortkey  string
	PageTitle string `"Go Short by Pummel :-)"`
	Message   string
}

var templates = template.Must(template.ParseGlob("resources/html/*"))

// simple html formular
// zeigt Eingabeformular
func MainHandler(storage storages.IStorage) http.Handler {

	var key string
	p := PageData{
		BaseUrl:   "http://localhost:8080",
		PageTitle: "Go Short by Pummel :-)",
	}

	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		// POST
		if url := r.PostFormValue("url"); url != "" {
			log.Print("save url: " + url)
			//w.Write([]byte(storage.Save(url)))
			key = storage.Save(url)

			p.Shortkey = key
		}

		log.Print("main after POST mit key: " + key)

		// Immer
		/*
			tmpl, err := template.ParseFiles("resources/html/main.html")
			if err != nil {
				log.Fatal("Fehler lesen Template:" + err.Error())
				p.Message = err.Error()
			}
			tmpl.Execute(w, p)
		*/
		// from https://golangcode.com/using-html-templates-from-a-folder-complied/
		err := templates.ExecuteTemplate(w, "main.html", p)
		if err != nil {
			log.Fatal("Cannot Get View ", err)
		}
	}

	return http.HandlerFunc(handleFunc)
}

// New URL
func EncodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		if url := r.PostFormValue("url"); url != "" {
			log.Print("save url: " + url)
			w.Write([]byte(storage.Save(url)))
			// TODO: Zeige Ergebnis und Eingabeformular
		}
	}

	return http.HandlerFunc(handleFunc)
}

// decodiert eine Kurzurl und ... ?
func DecodeHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/dec/"):]
		log.Print("read saved url for code: " + code)

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		w.Write([]byte(url))
	}

	return http.HandlerFunc(handleFunc)
}

// leitet auf echte URL weiter
func RedirectHandler(storage storages.IStorage) http.Handler {
	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/red/"):]
		log.Print("redirect to saved url for code: " + code)

		url, err := storage.Load(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("URL Not Found. Error: " + err.Error() + "\n"))
			return
		}

		http.Redirect(w, r, string(url), 301)
	}

	return http.HandlerFunc(handleFunc)
}
