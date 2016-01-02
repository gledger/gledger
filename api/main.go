package main

import (
	"flag"
	"net"
	"net/http"
	"os"

	"github.com/zombor/ledger"
)

func main() {
	listen := flag.String("listen", ":0",
		"TCP address (host:port) on which to listen for HTTP connections."+
			" Defaults to a random port."+
			" See http://golang.org/pkg/net/#Dial for examples.")
	journalPath := flag.String("journal", "", "File path to ledger journal file. Required")

	flag.Parse()

	if *journalPath == "" {
		println("-journal is a required flag")
		os.Exit(-1)
	}

	listener, err := net.Listen("tcp", *listen)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	println("listening on", listener.Addr().String())

	handler := restHandler{
		journal: ledger.NewJournal(
			ledger.NewFileReader(*journalPath),
		),
	}

	http.HandleFunc("/", handler.Root)
	http.Serve(listener, http.DefaultServeMux)
}
