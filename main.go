package main

import (
	"fmt"
	"net/http"

	"github.com/cneill/quote/pkg/prettify"
)

var listen = "localhost:9999"

func runServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/prettify", Prettify)
	mux.HandleFunc("/tags/reload", ReloadTags)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", Index)

	fmt.Printf("Listening on %s...\n", listen)

	if err := http.ListenAndServe(listen, mux); err != nil {
		return fmt.Errorf("listening error: %w", err)
	}

	return nil
}

func run() error {
	if err := prettify.LoadTags(); err != nil {
		return fmt.Errorf("run: failed to load tags: %w", err)
	}

	if err := runServer(); err != nil {
		return fmt.Errorf("run: failed running server: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
