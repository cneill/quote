package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/cneill/quote/pkg/prettify"
)

func doErr(w http.ResponseWriter, err error) {
	fmt.Printf("ERROR WITH REQUEST: %v\n", err)
	w.WriteHeader(http.StatusBadRequest)

	if _, err := w.Write([]byte(fmt.Sprintf("%v", err))); err != nil {
		fmt.Printf("error writing error: %v\n", err)
	}
}

// Index serves the main page.
func Index(w http.ResponseWriter, _ *http.Request) {
	t, err := template.New("index").ParseFiles("index.gohtml")
	if err != nil {
		doErr(w, fmt.Errorf("failed to load index template: %w", err))

		return
	}

	if err := t.Execute(w, nil); err != nil {
		doErr(w, fmt.Errorf("failed to execute index template: %w", err))

		return
	}
}

// Prettify serves the template containing the prettified text in a <textarea />.
func Prettify(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("prettify").ParseFiles("prettify.gohtml")
	if err != nil {
		doErr(w, fmt.Errorf("failed to load prettify template: %w", err))

		return
	}

	if err := req.ParseForm(); err != nil {
		doErr(w, fmt.Errorf("failed to parse form: %w", err))

		return
	}

	input := req.PostForm.Get("input")
	if input != "" {
		input = prettify.Prettify(input)
	}

	data := struct {
		Prettified string
	}{input}

	if err := t.Execute(w, data); err != nil {
		doErr(w, fmt.Errorf("failed to execute prettify template: %w", err))

		return
	}
}

// ReloadTags causes the tags to be re-populated from the tags.json file.
func ReloadTags(w http.ResponseWriter, req *http.Request) {
	if err := prettify.LoadTags(); err != nil {
		doErr(w, fmt.Errorf("failed to reload tags: %w", err))
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}
