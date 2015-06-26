package chikurin

import (
	"html/template"
	"net/http"

	"github.com/yosssi/ace"
)

type controllerData struct {
	Title  string
	Search bool
}

func (controller controllerData) render(w http.ResponseWriter, view string, data interface{}, funcs *template.FuncMap) {
	options := ace.Options{Asset: Asset}
	if funcs != nil {
		options.FuncMap = *funcs
	}

	tpl, err := ace.Load("view/layout", view, &options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
