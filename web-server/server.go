package main

import (
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, r.URL.Path[1:])
	// })

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8081", nil))
}
