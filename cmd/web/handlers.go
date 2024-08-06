package main

import (
	"html/template"
	"log"
	"net/http"
)

type structForTemplate struct { // this structure is intended to be passed to the template
	Name string
}

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interner Server Error ParseFiles", 500)
		return
	}

	if r.Method == "POST" {

		var totalNameString string

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		nameReceived := r.PostForm.Get("name") //	first use r.ParseForm()

		log.Printf("POST one line: %s\n", nameReceived)

		totalNameString = "Hello! " + nameReceived

		err = ts.Execute(w, structForTemplate{totalNameString})
		if err != nil {
			log.Println(err)
			http.Error(w, "Interner Server Error  MyTest", 500)
		}

	} else {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

		log.Println("Not POST")
	}
}
