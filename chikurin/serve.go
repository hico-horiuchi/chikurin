package chikurin

import (
	"net/http"

	"github.com/zenazn/goji"
)

func Serve() {
	goji.Get("/assets/*", http.StripPrefix("/assets", http.FileServer(http.Dir("assets"))))
	goji.Get("/:datacenter/:client", statusPage)
	goji.Serve()
}
