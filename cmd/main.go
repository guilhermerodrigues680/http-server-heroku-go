// Package main implementa metodos para responder a alguma requisicao HTTP
//
//
package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr)
	fmt.Fprint(w, "Ola direto do heroku")
}

func main() {

	addr := ":8082"
	http.HandleFunc("/", handler)
	log.Printf("Servidor iniciado. IP -> %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
