package chikurin

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

type datacentersData struct {
	controllerData
	Datacenters []datacenterStruct
}

func datacentersController(c web.C, w http.ResponseWriter, r *http.Request) {
	data := datacentersData{
		controllerData: controllerData{
			Title:  "Datacenters",
			Search: true,
		},
		Datacenters: config.Datacenters,
	}

	data.render(w, "view/datacenters", data, nil)
}
