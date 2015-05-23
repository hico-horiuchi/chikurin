package chikurin

import (
	"net/http"

	"github.com/zenazn/goji"
)

func Serve() {
	goji.Get("/assets/*", http.FileServer(AssetFileSystem{}))

	goji.Get("/:datacenter/:client", statusController)
	if config.ShowDatacenters || config.ShowClients {
		goji.Get("/:datacenter", clientsController)
	}
	if config.ShowDatacenters {
		goji.Get("/", datacenterController)
	}

	goji.Serve()
}
