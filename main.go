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
	path := "/static"
	fs := http.FileServer(http.Dir(path))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Home)
	http.HandleFunc("/result", result)
	fmt.Printf("Fetching server...")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if r.URL.Path != "/result" {
			tmpl.ExecuteTemplate(w, "404.html", nil)
			return
		}
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "static/index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			tmpl.ExecuteTemplate(w, "400.html", nil)
			return
		} else {
			_, _ = w.Write([]byte("Submission recorded!"))
		}
	default:
		tmpl.ExecuteTemplate(w, "500.html", nil)
		return
	}

}

func result(w http.ResponseWriter, r *http.Request) {
	data := a.Ascii(r.FormValue("text"))
	tmpl.ExecuteTemplate(w, "result.html", data)
}
