package chikurin

import "encoding/json"

type eventStruct struct {
	Id          string
	Client      clientStruct
	Check       checkStruct
	Occurrences int
	Action      string
}

func (dc datacenterStruct) getEventsClient(name string) ([]eventStruct, error) {
	var events []eventStruct

	request, err := dc.makeRequest("GET", "/events/"+name, nil)
	if err != nil {
		return events, err
	}

	contents, err := sensuAPI(request)
	if err != nil {
		return events, err
	}

	err = json.Unmarshal(contents, &events)
	return events, err
}
