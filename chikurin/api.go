package chikurin

import (
	"io/ioutil"
	"net/http"
)

func sensuAPI(request *http.Request) ([]byte, error) {
	var body []byte

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return body, err
	}

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return body, err
	}

	defer response.Body.Close()
	return body, nil
}
