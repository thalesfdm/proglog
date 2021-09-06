package main

import (
	"github.com/thalesfdm/proglog/internal/server"
	"log"
)

func main() {
	s := server.NewHTTPServer(":8080")
	log.Fatal(s.ListenAndServe())
}
