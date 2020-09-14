package main

import (
	"net/http"

	"github.com/labstack/gommon/log"
)

func main() {

	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
