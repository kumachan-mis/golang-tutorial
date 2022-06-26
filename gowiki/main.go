package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", MakeHandler(ViewHandler))
	http.HandleFunc("/edit/", MakeHandler(EditHandler))
	http.HandleFunc("/save/", MakeHandler(SaveHandler))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
