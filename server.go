package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/intervention-engine/fhir/server"
)

func main() {
	host := os.Getenv("MONGODB_PORT_27017_TCP_ADDR")
	port := os.Getenv("MONGODB_PORT_27017_TCP_PORT")

	fmt.Printf("Connecting to %s:%s", host, port)
	s := server.NewServer(host)

	s.Run()
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "FHIR Server Yay! \\o/")
}
