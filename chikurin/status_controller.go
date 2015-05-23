package chikurin

import (
	"net/http"

	"github.com/yosssi/ace"
	"github.com/zenazn/goji/web"
)

type statusData struct {
	layoutData
	Timestamp string
	Client    clientStruct
	Events    []eventStruct
}

func statusController(c web.C, w http.ResponseWriter, r *http.Request) {
	var data statusData

	dc, err := config.selectDatacenter(c.URLParams["datacenter"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Client, err = dc.getClientsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Events, err = dc.getEventsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl, err := ace.Load("view/layout", "view/status", &ace.Options{
		Asset: Asset,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.Title = data.Client.Name
	err = tpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
