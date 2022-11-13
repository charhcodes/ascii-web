package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", serverHandler)
	server := http.FileServer(http.Dir("./templates"))
	http.Handle("/template/", http.StripPrefix("/template/", server)) //stripprefix cuts off static and takes anything after as a request
	fmt.Println("Listening for port 8080...")
	http.ListenAndServe(":8080", nil)
}

func serverHandler(res http.ResponseWriter, req *http.Request) {
	//template := template.Must(template.ParseGlob("static/templates/temp.html"))
	if req.URL.Path != "/" { //this is where the path ends
		http.Error(res, "404 not found.", http.StatusNotFound)
		return
	}
	switch req.Method {
	case "GET":
		http.ServeFile(res, req, "templates/template.html")
	case "POST":
		if err := req.ParseForm(); err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest) // if there is an error it returns bad request 400
			return
		} else {
			fmt.Fprintf(res, "it worked!")
		}
		//input := req.FormValue("input")
	default:
		//method not supported
		http.Error(res, "Method is not supported.", http.StatusUnsupportedMediaType)
		return
	}

}
