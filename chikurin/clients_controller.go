package chikurin

import (
	"net/http"

	"github.com/hico-horiuchi/ohgi/sensu"
	"github.com/zenazn/goji/web"
)

type clientsData struct {
	controllerData
	Datacenter *datacenterStruct
	Clients    []sensu.ClientStruct
}

func clientsController(c web.C, w http.ResponseWriter, r *http.Request) {
	datacenter, err := config.selectDatacenter(c.URLParams["datacenter"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := clientsData{
		controllerData: controllerData{
			Title:  datacenter.Name,
			Search: true,
		},
		Datacenter: datacenter,
	}

	data.Clients, err = datacenter.sensuAPI().GetClients(-1, -1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data.render(w, "view/clients", data, nil)
}
