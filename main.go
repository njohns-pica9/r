package main

// Open url in browser:
// http://localhost:14000/app

import (
	// "github.com/RangelReale/osin"
	"github.com/njohns-pica9/cas/storage"
	//	"log"
	//	"net/http"
	//	"net/url"
)

var rstore *storage.Restorage

func main() {
	rstore = storage.NewStorage()

	initServer()

	rstore.Close()
}
