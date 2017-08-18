<<<<<<< HEAD
package controller
=======
package controller //entry point of controller layer
>>>>>>> master

import (
	"html/template"
	"net/http"
)

<<<<<<< HEAD
=======
//populate homeTemplate and create the controller
>>>>>>> master
var (
	homeController home
	shopController shop
)

<<<<<<< HEAD
func Startup(templates map[string]*template.Template) {
=======
func Startup(templates map[string]*template.Template) { //startup function
>>>>>>> master
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
<<<<<<< HEAD
	http.Handle("/img/", http.FileServer(http.Dir("public")))
=======
	http.Handle("/img/", http.FileServer(http.Dir("public"))) //static resources from main.go
>>>>>>> master
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
