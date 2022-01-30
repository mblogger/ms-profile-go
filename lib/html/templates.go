package templates

import (
	"biodata/model"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Read(templateName string, templateToRender string, data model.HomeData, context *gin.Context) {
	tpl := templateToRender
	t, err := template.New(templateName).Parse(tpl)
	if err != nil {
		log.Panic("Unable to parse template!")
	}

	t.Execute(context.Writer, data)
}

func ReadFile(templateName string, fileName string) {
	file := readFile(fileName)
	t, err := template.New(templateName).ParseFiles(file.Name())
	if err != nil {
		log.Panic("Unable to parse files at the location")
	}
	data := struct {
		Title string
	}{
		Title: "MS - Profile",
	}

	t.Execute(os.Stdout, data)
}

func readFile(fileName string) os.File {
	f, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error reading file")
	}
	defer f.Close()

	_, err = fmt.Fprintf(os.Stdout, "%s:", f.Name())

	if err != nil {
		log.Panic("Unable to print output of file!")
	}

	return *f
}
