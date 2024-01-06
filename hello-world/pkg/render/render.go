package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplates(w http.ResponseWriter, tmp string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmp)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("err parsing template:", err)
		return
	}
}
