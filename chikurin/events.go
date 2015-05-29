package chikurin

import (
	"encoding/json"
	"time"
)

type eventStruct struct {
	Id          string
	Client      clientStruct
	Check       checkStruct
	Occurrences int
	Action      string
	Since       string
	At          string
}

func (dc datacenterStruct) getEventsClient(name string) ([]eventStruct, error) {
	var events []eventStruct
	var elapsed time.Duration
	var then time.Time

	request, err := dc.makeRequest("GET", "/events/"+name)
	if err != nil {
		return events, err
	}

	contents, err := sensuAPI(request)
	if err != nil {
		return events, err
	}

	err = json.Unmarshal(contents, &events)
	if err != nil {
		return events, err
	}

	for i := range events {
		elapsed = time.Duration(events[i].Occurrences) * time.Duration(events[i].Check.Interval) * time.Second
		then = time.Unix(events[i].Client.Timestamp, 0)
		events[i].Since = then.Add(-1 * elapsed).Format("2006/01/02 15:04:05")
		events[i].At = utoa(events[i].Client.Timestamp)
	}

	return events, nil
}
