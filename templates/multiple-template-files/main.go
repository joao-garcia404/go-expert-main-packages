package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(
		template.New("content.html").
			Funcs(template.FuncMap{"ToUpper": ToUpper}). // Mapping function to be used at HTML file
			ParseFiles(templates...),
	)

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"Java", 30},
		{"Python", 20},
	})

	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
