package main

import (
	"log"
	"net/http"
	"os"

	handlers "github.com/krisztiking/microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "prodict-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}
