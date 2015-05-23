package chikurin

import (
	"net/http"

	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
)

type clientsData struct {
	layoutData
	Datacenter datacenterStruct
	Clients    []clientStruct
}

func clientsController(c web.C, w http.ResponseWriter, r *http.Request) {
	var data clientsData

	dc, err := config.selectDatacenter(c.URLParams["datacenter"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Clients, err = dc.getClients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := ace.Load("view/layout", "view/clients", &ace.Options{
		Asset: Asset,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Title = dc.Name
	data.Search = true
	data.Datacenter = dc
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
