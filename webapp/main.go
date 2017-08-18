package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/jamaykajones/lss/webapp/controller"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template { //maps of strings to templates
	result := make(map[string]*template.Template)
	const basePath = "D:/Users/jjones/Go/src/github.com/jamaykajones/lss/templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))             // loading in the template
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html")) //loding subtemps the layout will use
	dir, err := os.Open(basePath + "/content")                                           //load in conetent DIR
	if err != nil {
		panic("Failed to open template block directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1) //read file
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fileInfo := range fis {
		file, err := os.Open(basePath + "/content/" + fileInfo.Name()) //open each file
		if err != nil {
			panic("Failed to open template '" + fileInfo.Name() + "'")
		}
		content, err := ioutil.ReadAll(file) // read each file
		if err != nil {
			panic("Failed to read content from file '" + fileInfo.Name() + "'")
		}
		file.Close()                          // close the file after it has been read
		tmpl := template.Must(layout.Clone()) // create template instance by cloning
		_, err = tmpl.Parse(string(content))  // parse the contents we just read from the tmpl file
		if err != nil {
			panic("Failed to parse contents of '" + fileInfo.Name() + "' as template")
		}
		result[fileInfo.Name()] = tmpl //add file to result map
	}
	return result
}
