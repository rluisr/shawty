package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/rluisr/shawty/handlers"
	"github.com/rluisr/shawty/storages"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	storage := &storages.Redis{}
	err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handlers.EncodeHandler(storage))
	http.Handle("/dec/", handlers.DecodeHandler(storage))
	http.Handle("/red/", handlers.RedirectHandler(storage))

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
