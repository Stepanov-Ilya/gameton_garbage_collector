package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gameton/Objects"
	"io/ioutil"
	"net/http"
)

func PlayerUniversal() Objects.UniverseResponse {

	client := http.Client{}
	request, err := http.NewRequest("GET", URL+"/player/universe", bytes.NewBuffer(nil))
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Auth-Token", Token)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	universal := Objects.UniverseResponse{}

	err = json.Unmarshal(body, &universal)
	if err != nil {
		fmt.Println(err)
	}

	return universal
}
