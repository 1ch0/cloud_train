package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", myHandler)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	http.ListenAndServe(":8080", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client IP:%s,code:%s", r.RemoteAddr, "200")
	for k, v := range r.Header {
		w.Header().Set(k, "")
		for _, v1 := range v {
			w.Header().Add(k, v1)
		}
	}
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	w.WriteHeader(200)
}
