package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	basePath := flag.String("b", "/", "base path")
	flag.Parse()

	http.Handle(*basePath, http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on http://0.0.0.0:%s%s\n", *directory, *port, *basePath)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
