package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var dir = flag.String("dir", ".", "path where the files to serve will be fetched")
var lnet = flag.String("lnet", "unix", "server listening network")
var laddr = flag.String("laddr", "fs.sock", "server listening address (host:port) or socket file")

func main() {
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ln, err := new(net.ListenConfig).Listen(ctx, *lnet, *laddr)
	if err != nil {
		log.Fatal(err)
	}
	h := http.FileServer(http.Dir(*dir))
	srv := http.Server{Handler: h}

	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
	go func() {
		log.Printf("**canceling** after signal: %v\n", <-sigch)
		cancel()
		srv.Close()
	}()

	log.Printf("server listening on %v://%v", *lnet, *laddr)
	err = srv.Serve(ln)
	if *lnet == "unix" {
		os.Remove(*laddr)
	}
	if err != nil {
		log.Fatal(err)
	}
}
