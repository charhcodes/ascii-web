package main

import (
	"asciiart-web/asciifiles"
	"net/http"
	"strings"
	"text/template"
)

type Text struct {
	Normaltext string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("html/*.html"))
	// http.HandleFunc("/save", save)

}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/process", Processor)
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/Error404", Error404)
	http.HandleFunc("/Error400", Error500)
	http.HandleFunc("/Error500", Error400)
}

// If the URL path is not "/", then serve the 404 page. Otherwise, serve the index page
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		//w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "html/Error404.html")
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)

}

func Processor(w http.ResponseWriter, r *http.Request) {

	fontval := r.FormValue("usersinput")
	splitfont := strings.Split(fontval, "")

	if fontval != "standard" && fontval != "shadow" && fontval != "thinker-toy" {
		//w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "html/Error500.html")
		return
	}
	textval := r.FormValue("asciitext")

	splittxt := strings.Split(textval, "")

	text := asciifiles.Asciiart(splittxt, splitfont)

	data := Text{
		Normaltext: text,
	}
	for _, word := range splittxt {
		for _, character := range word {
			if character > 128 {
				w.WriteHeader(http.StatusNotFound)
				http.ServeFile(w, r, "html/Error400.html")
				return
			}
		}
	}
	tpl.ExecuteTemplate(w, "index.html", data)

	//fmt.Print(asciifiles.Asciiart(textval, fontval))

}

// The function Error404 takes in a response writer and a request, and then executes the Error404.html
// template
func Error404(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Error404.html", nil)
}

func Error500(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Error500.html", nil)
}

func Error400(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Error400.html", nil)
}
