package chikurin

import (
	"net/http"
	"strconv"
)

type datacenterStruct struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
}

func (dc datacenterStruct) makeRequest(method string, namespace string) (*http.Request, error) {
	url := "http://" + dc.Host + ":" + strconv.Itoa(dc.Port) + namespace
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	if dc.User != "" && dc.Password != "" {
		request.SetBasicAuth(dc.User, dc.Password)
	}

	return request, nil
}
