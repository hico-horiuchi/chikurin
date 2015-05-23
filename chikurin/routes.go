package chikurin

import (
	"net/http"

	"github.com/zenazn/goji"
)

func Serve() {
	goji.Get("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))

	goji.Get("/:datacenter/:client", statusController)
	if config.ShowDatacenters || config.ShowClients {
		goji.Get("/:datacenter", clientsController)
	}
	if config.ShowDatacenters {
		goji.Get("/", datacenterController)
	}

	goji.Serve()
}
