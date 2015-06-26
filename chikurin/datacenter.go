package chikurin

import "github.com/hico-horiuchi/ohgi/sensu"

type datacenterStruct struct {
	sensu.API
	Name string `json:"name"`
}

func (datacenter datacenterStruct) sensuAPI() *sensu.API {
	return &sensu.API{
		Host:     datacenter.Host,
		Port:     datacenter.Port,
		User:     datacenter.User,
		Password: datacenter.Password,
	}
}
