package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/krisztiking/microservices"
)

func main() {
	l := log.New(os.Stdout, "prodict-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil)
}
