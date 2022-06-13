package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Inicio)
	log.Println("Servidor corriendo...")

}

func Inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Holaaaaa")

}
