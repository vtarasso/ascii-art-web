package datafile

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Word struct {
	InputWord  string
	Font       string
	OutputWord string
}

type ErrorStruct struct {
	Status       int
	ErrorMessage string
}

func Handler() {
	mux := http.NewServeMux()
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	mux.HandleFunc("/", firstHandler)
	mux.HandleFunc("/ascii-art", postHandler)
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal("ERROR:Server not listening")
	}
}

func firstHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		errorHandler(w, 404) // 404 Not Found
		return
	}
	if r.URL.Path == "/" && r.Method == "GET" {
		tmpl, err := template.ParseFiles("assets/templates/index.html")
		if err != nil {
			w.WriteHeader(500)
			errorHandler(w, 500)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			w.WriteHeader(500)
			errorHandler(w, 500)
			return
		}
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		errorHandler(w, 405) // 405 Method Not Allowed
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	res, status := Asciiart(text, banner)

	if text == "" || banner == "" {
		w.WriteHeader(500)
		errorHandler(w, 500) // 500 Internal Server Error
		return
	}

	if status == 500 {
		w.WriteHeader(500)
		errorHandler(w, 500)
		return
	} else if status == 400 {
		w.WriteHeader(400)
		errorHandler(w, 400) // 400 Bad Request
		return
	}
	word := Word{text, banner, res}
	tmpl, err := template.ParseFiles("assets/templates/index.html")
	if err != nil {
		w.WriteHeader(500)
		errorHandler(w, 500)
	}
	err = tmpl.Execute(w, word)
	if err != nil {
		w.WriteHeader(500)
		errorHandler(w, 500)
		return
	}
}

func errorHandler(w http.ResponseWriter, status int) {
	Res := ErrorStruct{
		Status:       status,
		ErrorMessage: http.StatusText(status),
	}

	tmpl, err := template.ParseFiles("assets/templates/error.html")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Error 500")
		return
	}
	err = tmpl.Execute(w, Res)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintln(w, "Internal Server Error 500")
		return
	}
}
