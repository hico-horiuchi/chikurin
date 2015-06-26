package chikurin

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"
)

var config configStruct
var timeout = 3 * time.Second

type configStruct struct {
	Datacenters     []datacenterStruct
	ShowDatacenters bool   `json:"show_datacenters"`
	ShowClients     bool   `json:"show_clients"`
	Bind            string `json:"bind"`
	Log             string `json:"log"`
}

func LoadConfig() {
	bytes, err := ioutil.ReadFile("/etc/chikurin.json")
	checkError(err)

	err = json.Unmarshal(bytes, &config)
	checkError(err)
}

func (config configStruct) selectDatacenter(name string) (*datacenterStruct, error) {
	switch {
	case len(config.Datacenters) == 0:
		return nil, errors.New("chikurin: no datacenters in config")
	case name == "":
		return &config.Datacenters[0], nil
	}

	for _, datacenter := range config.Datacenters {
		if datacenter.Name == name {
			return &datacenter, nil
		}
	}

	return nil, errors.New("chikurin: no such datacenter in config")
}
