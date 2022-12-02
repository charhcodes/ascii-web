package main

import (
	a "ascii-web-practice2/ascii-art"
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("static/*.html"))
}

func main() {
	path := "static"
	fs := http.FileServer(http.Dir(path))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Home)
	http.HandleFunc("/result", result)
	http.HandleFunc("/404", Err404)
	http.HandleFunc("/400", Err500)
	http.HandleFunc("/500", Err400)
	fmt.Printf("Fetching server...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl.ExecuteTemplate(w, "404.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func result(w http.ResponseWriter, r *http.Request) {
	bannerop := r.FormValue("font")  //for the banner choice
	textinput := r.FormValue("text") // for the text input

	// check if characters are applicable
	runes := []rune(textinput)
	for i := range runes {
		if runes[i] > 127 {
			//w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, "static/400.html")
			return
		}
	}

	text := a.Ascii(bannerop, textinput) //runs our ascii-code over the banner and text
	tmpl.ExecuteTemplate(w, "result.html", text)
}

func Err404(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "404.html", nil)
}

func Err500(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "500.html", nil)
}

func Err400(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "400.html", nil)
}
