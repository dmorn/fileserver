package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var dir = flag.String("dir", ".", "path where the files to serve will be fetched")
var port = flag.String("port", "8080", "server listening port")

func main() {
	flag.Parse()

	fs := http.Dir(*dir) // files in local directory
	p := ":" + *port

	log.Printf("%s listening on %s", os.Args[0], p)
	if err := http.ListenAndServe(p, FileServer(fs)); err != nil {
		fmt.Printf("%s stopping, reason: %v\n", os.Args[0], err)
	}
}

type fileHandler struct {
	handler http.Handler
}

func FileServer(root http.FileSystem) http.Handler {
	return &fileHandler{http.FileServer(root)}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%v] %v", r.Method, r.URL.Path)
	f.handler.ServeHTTP(w, r)
}
