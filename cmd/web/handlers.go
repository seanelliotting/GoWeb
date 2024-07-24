package main

import (
	"html/template"
	"log"
	"net/http"
)

type NameStruct struct {
	Name string
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		var NameString string

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		Name := r.PostForm.Get("name") //	first use r.ParseForm()

		//Name := r.FormValue("name")		//	another method of extracting data without r.ParseForm()

		log.Println("POST")
		log.Println(Name)

		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Interner Server Error", 500)
			return
		}

		NameString = "Hello! " + Name

		ts.Execute(w, NameStruct{NameString})

	} else {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
		}

		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Interner Server Error", 500)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		log.Println("Not POST")

		//w.Write([]byte("Hello from GoWeb!"))		// display text in the browser window
	}
}
