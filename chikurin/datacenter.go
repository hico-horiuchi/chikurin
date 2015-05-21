package chikurin

import (
	"io"
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

func (dc datacenterStruct) makeRequest(method string, namespace string, payload io.Reader) (*http.Request, error) {
	url := "http://" + dc.Host + ":" + strconv.Itoa(dc.Port) + namespace
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	if dc.User != "" && dc.Password != "" {
		request.SetBasicAuth(dc.User, dc.Password)
	}

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}
