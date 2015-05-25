package chikurin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var config configStruct
var timeout = 3 * time.Second

type configStruct struct {
	Datacenters     []datacenterStruct
	ShowDatacenters bool `json:"show_datacenters"`
	ShowClients     bool `json:"show_clients"`
	Timeout         int
	Bind            string
	Log             string
}

func LoadConfig() {
	bytes, err := ioutil.ReadFile("/etc/chikurin.json")
	checkError(err)

	json.Unmarshal(bytes, &config)
	checkError(err)

	if config.Timeout > 0 {
		timeout = time.Duration(config.Timeout) * time.Second
	}
	http.DefaultClient.Timeout = timeout
}

func (c configStruct) selectDatacenter(name string) (datacenterStruct, error) {
	switch {
	case len(c.Datacenters) < 1:
		return datacenterStruct{}, errors.New("chikurin: no datacenters in config")
	case name == "":
		return c.Datacenters[0], nil
	}

	for _, dc := range c.Datacenters {
		if dc.Name == name {
			return dc, nil
		}
	}

	return datacenterStruct{}, errors.New("chikurin: no such datacenter in config")
}
