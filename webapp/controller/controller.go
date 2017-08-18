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
	homeController.standLocatorTemplate = templates["stand_locator.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	shopController.productTemplate = templates["shop_detail.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public"))) //static resources from main.go
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
