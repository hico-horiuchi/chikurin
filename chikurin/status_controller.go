package chikurin

import (
	"html/template"
	"net/http"

	"github.com/hico-horiuchi/ohgi/sensu"
	"github.com/zenazn/goji/web"
)

type statusData struct {
	controllerData
	Client *sensu.ClientStruct
	Events []sensu.EventStruct
}

func statusController(c web.C, w http.ResponseWriter, r *http.Request) {
	datacenter, err := config.selectDatacenter(c.URLParams["datacenter"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client, err := datacenter.sensuAPI().GetClientsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := statusData{
		controllerData: controllerData{
			Title:  client.Name,
			Search: false,
		},
		Client: &client,
	}

	data.Events, err = datacenter.sensuAPI().GetEventsClient(c.URLParams["client"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.render(w, "view/status", data, &template.FuncMap{
		"at":    at,
		"since": since,
	})
}
