package main

import (
	"html/template"
	"os"
)

type content struct {
	Title string
	Text  string
}

func main() {
	template := template.Must(template.ParseFiles(
		"partial.html",
		"template.html",
	))

	content := &content{
		Title: "Sed Ut Perspiciatis",
		Text:  "Lorem ipsum dolor sit amet, consectetur adipiscing.",
	}

	if err := template.ExecuteTemplate(os.Stdout, "MAIN", content); err != nil {
		panic(err)
	}
}
