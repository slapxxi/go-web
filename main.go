package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

const STATIC_DIR = "/public"
const STATIC_PREFIX = "/static/"

func main() {
	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir(getPublicDir()))

	mux.Handle(STATIC_PREFIX, http.StripPrefix(STATIC_PREFIX, fileHandler))
	mux.HandleFunc("/", indexHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", []byte{})
}

func getPublicDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, STATIC_DIR)
}
