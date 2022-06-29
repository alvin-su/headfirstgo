package main

import (
	"html/template"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func excuteTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)
}

type Part struct {
	Name  string
	Count int
}

func main() {
	//text := "Here's my template!\n"

	/*templateText := "Template start\nAction:{{.}}\nTemplate end\n"

	tmpl, err := template.New("test").Parse(templateText)
	check(err)
	err = tmpl.Execute(os.Stdout, "ABC")
	check(err)*/

	excuteTemplate("Dot is:{{.}}!\n", "ABC")
	excuteTemplate("Dot is:{{.}}!\n", 123)

	excuteTemplate("start {{if .}}Dot is true!{{end}} finish!\n", true)
	excuteTemplate("start {{if .}}Dot is true!{{end}} finish!\n", false)

	templateText := "Before loop:{{.}}\n{{range .}}In loop {{.}}\n{{end}}After loop:{{.}}"
	excuteTemplate(templateText, []string{"do", "re", "mi"})

	templateText = "Prices:\n{{range .}}${{.}}\n{{end}}"
	excuteTemplate(templateText, []float64{1.25, 0.99, 27})

	templateText = "Name:{{.Name}}\nCount:{{.Count}}\n"
	excuteTemplate(templateText, Part{Name: "Fuses", Count: 5})
	excuteTemplate(templateText, Part{Name: "Cables", Count: 2})
}
