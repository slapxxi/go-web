package main

import (
	"fmt"
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
	mux.HandleFunc("POST /threads/{id}", threadsHandler)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServeTLS("cert.pem", "key.pem")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving index")
	files := []string{
		"templates/layout.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", []byte{})
}

func threadsHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte(id))
}

func getPublicDir() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(wd, STATIC_DIR)
}
