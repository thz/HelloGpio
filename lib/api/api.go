package api

import (
	"bytes"
	"net/http"
)

const (
	API_VAL_UP   = "0"
	API_VAL_DOWN = "1"
	API_VAL_STOP = "1"
)

func Apicall(uri string, data string) error {

	// data goes into request body
	buf := bytes.NewBuffer([]byte(data))
	request, err := http.NewRequest("POST", "http://api.iot.example.com:8002"+uri, buf)
	if err != nil {
		return err
	}

	request.Header.Add("X-Apikey", `xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx`)
	client := &http.Client{}
	_, err2 := client.Do(request) // ignore response

	return err2
}
