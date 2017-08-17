package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/jamaykajones/lss/webapp/model"
	"github.com/jamaykajones/lss/webapp/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (s shop) registerRoutes() {
	http.HandleFunc("/shop", s.handleShop)
	http.HandleFunc("/shop/", s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`) // regex looking for a path that starts with /shop/
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1]) //converting string to a int
		s.handleCategory(w, r, categoryID)        // pass in categoryID
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		s.shopTemplate.Execute(w, vm)
	}
}

//populate viewmodel
func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	vm := viewmodel.NewShopDetail(products)
	s.categoryTemplate.Execute(w, vm)
}
