package chikurin

import (
	"net/http"

	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
)

type statusStruct struct {
	Title  string
	Client clientStruct
	Events []eventStruct
}

func statusPage(c web.C, w http.ResponseWriter, r *http.Request) {
	var status statusStruct

	dc, err := config.selectDatacenter(c.URLParams["datacenter"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status.Client, err = dc.getClientsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status.Events, err = dc.getEventsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := ace.Load("view/layout", "view/status", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	status.Title = status.Client.Name
	err = tpl.Execute(w, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
