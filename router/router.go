package router

import (
	"log"
	"net/http"
	"time"
	"websites-status-checker/internal/handlers"
	"websites-status-checker/internal/service"
)

func HandleRequest() {
	go func() {
		for {
			service.Traverse()
			time.Sleep(20 * time.Second)
		}
	}()
	http.HandleFunc("/post", handlers.PostHandler)
	http.HandleFunc("/get", handlers.GetHandler)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
