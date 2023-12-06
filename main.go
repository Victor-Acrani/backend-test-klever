package main

import (
	"flag"
	"log"
	"os"

	"github.com/Victor-Acrani/backend-test-klever/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "server address")
	flag.Parse()

	s := api.NewServer(*listenAddr)
	err := s.Start()
	if err != nil {
		log.Println("s.Start() ", err.Error())
		os.Exit(1)
	}
}
