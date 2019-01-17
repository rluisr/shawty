package main

import (
	"log"
	"net/http"
	"os"

	"github.com/didip/shawty/handlers"
	"github.com/didip/shawty/storages"
)

func main() {
	// this should be read from config!
	credentials := storages.DbCredentials{
		Host: "localhost",
		Port: 6379,
		Name: 0,
		User: "",
		Pass: "",
	}

	storage := &storages.Redis{}
	err := storage.Init(credentials)
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handlers.MainHandler(storage))
	http.Handle("/enc", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

	// port should be read from config!
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	println("running url shortener on http://localhost:" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
