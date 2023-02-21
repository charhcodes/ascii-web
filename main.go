package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	http.HandleFunc("/result", execute)
	// download
	http.HandleFunc("/download", execute)

	http.HandleFunc("/404", Err404)
	http.HandleFunc("/400", Err400)
	http.HandleFunc("/500", Err500)
	fmt.Printf("Fetching server...")
	http.ListenAndServe(":8080", nil)
}

func Ascii(str string) string {
	split := strings.Split(str, `\n`)

	text, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("cannot read text file")
		log.Fatal(err)
	}

	output := ""
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(split); i++ {
		if string(split[i]) == "" {
			fmt.Println()
		} else {
			for j := 0; j < 8; j++ {
				for k := 0; k < len(string(split[i])); k++ {
					output += lines[int(((rune(split[i][k])-32)*9+1))+j]
				}
				output += "\n"
			}
		}
	}
	return output
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl.ExecuteTemplate(w, "404.html", nil)
		return
	}
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func result(textinput string) (string, error) {
	// check if characters are applicable
	runes := []rune(textinput)
	for i := range runes {
		if runes[i] > 127 {
			fmt.Println("cannot print specified characteers")
			return "", errors.New("cannot print specified characteers")
		}
	}
	text := Ascii(textinput)
	return text, nil
}

func execute(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	submit := r.FormValue("submit")
	download := r.FormValue("download")

	textinput, err := result(text)

	if len(submit) != 0 {
		if err != nil {
			tmpl.ExecuteTemplate(w, "500.html", nil)
		}
		if len(textinput) == 0 || textinput == " " {
			fmt.Println("blank error")
			tmpl.ExecuteTemplate(w, "500.html", nil)
		}

		tmpl.ExecuteTemplate(w, "result.html", textinput)
	}

	if len(download) != 0 {
		f, _ := os.Open("download.txt")
		defer f.Close()

		w.Header().Set("Content-Disposition", "attachment; filename=download.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(textinput)))

		_, err := w.Write([]byte(textinput))
		if err != nil {
			fmt.Println("error writing response:", err)
		}
	}

}

// need to copy over user input text box stuff
// func downloadHandler(w http.ResponseWriter, r *http.Request) {
// 	textinput := r.FormValue("text")
// 	fmt.Println(textinput)
// 	fmt.Println("textinput:", textinput)
// 	if len(textinput) == 0 || textinput == " " {
// 		fmt.Println("blank error download func")
// 		tmpl.ExecuteTemplate(w, "500.html", nil)
// 	}
// 	text := Ascii(textinput)

// 	f, _ := os.Open("download.txt")
// 	defer f.Close()

// 	w.Header().Set("Content-Disposition", "attachment; filename=download.txt")
// 	w.Header().Set("Content-Type", "text/plain")
// 	w.Header().Set("Content-Length", strconv.Itoa(len(text)))

// 	_, err := w.Write([]byte(text))
// 	if err != nil {
// 		fmt.Println("error writing response:", err)
// 	}
// }

func Err404(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "404.html", nil)
}

func Err500(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "500.html", nil)
}

func Err400(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "400.html", nil)
}
