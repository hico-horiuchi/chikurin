package chikurin

import "encoding/json"

type clientStruct struct {
	Name          string
	Address       string
	Subscriptions []string
	Timestamp     int64
	At            string
}

func (dc datacenterStruct) getClientsClient(name string) (clientStruct, error) {
	var client clientStruct

	request, err := dc.makeRequest("GET", "/clients/"+name, nil)
	if err != nil {
		return client, err
	}

	contents, err := sensuAPI(request)
	if err != nil {
		return client, err
	}

	err = json.Unmarshal(contents, &client)
	if err != nil {
		return client, err
	}

	client.At = utoa(client.Timestamp)
	return client, err
}
