package controller //entry point of controller layer

import (
	"html/template"
	"net/http"
)

//populate homeTemplate and create the controller
var (
	homeController home
	shopController shop
)

func Startup(templates map[string]*template.Template) { //startup function
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public"))) //static resources from main.go \/
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
