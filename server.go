package hiramekishare

import (
	"html/template"
	"log"
	"net/http"
)

var (
	templates *template.Template
	err       error
)

func StartApplication() {
	var templateFiles = []string{
		"",
	}

	templates, err = template.ParseFiles(templateFiles...)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", nil)
	http.ListenAndServe(":8080", nil)
}
