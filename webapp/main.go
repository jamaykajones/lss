package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jamayka.jones/lss/webapp/viewmodel"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates[requestedFile+".html"]
		var context interface{} //empty interface
		switch requestedFile {
		case "shop":
			context = viewmodel.NewShop()
		default:
			context = viewmodel.NewHome()
		}
		if template != nil {
			err := template.Execute(w, context)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(404) //statusNotFound
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
}

func populateTemplates() map[string]*template.Template { //maps of strings to templates
	result := make(map[string]*template.Template)
	const basePath = "templates"
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
