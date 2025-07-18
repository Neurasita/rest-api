package main

import (
	"net/http"

	"github.com/Neurasita/rest-api/internal/server"
)

func main() {
	mux := http.NewServeMux()
	var h http.Handler = mux

	s := server.New(h)

	s.Start()
	s.WaitShutdown()

}
