package chikurin

import (
	"net/http"

	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
)

type datacenterData struct {
	layoutData
	Datacenters []datacenterStruct
}

func datacenterController(c web.C, w http.ResponseWriter, r *http.Request) {
	var data datacenterData

	tpl, err := ace.Load("view/layout", "view/datacenter", &ace.Options{
		Asset: Asset,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Title = "Datacenters"
	data.Search = true
	data.Datacenters = config.Datacenters
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
