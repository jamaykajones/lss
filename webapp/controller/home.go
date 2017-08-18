package controller

import (
	"html/template"
	"net/http"

	"github.com/jamaykajones/lss/webapp/viewmodel"
)

type home struct {
<<<<<<< HEAD
	homeTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
=======
	homeTemplate *template.Template //hold config data from home page template
}

func (h home) registerRoutes() {
	http.HandleFunc("/home", h.handleHome) //Startup func creates routehandlers then call this method
	http.HandleFunc("/", h.handleHome)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) { //responsibility of processing request
>>>>>>> master
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}
