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

func PlayerTravel(planetsRoad []string) Objects.TravelResponse {
	postBody, _ := json.Marshal(Objects.TravelRequest{Planets: planetsRoad})

	client := http.Client{}

	request, err := http.NewRequest("POST", URL+"/player/travel", bytes.NewBuffer(postBody))
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

	travel := Objects.TravelResponse{}

	err = json.Unmarshal(body, &travel)
	if err != nil {
		fmt.Println(err)
	}

	return travel
}

func PlayerCollect(garbage map[string][]interface{}) Objects.ResponseCollect {
	postBody, _ := json.Marshal(garbage)

	client := http.Client{}

	request, err := http.NewRequest("POST", URL+"/player/collect", bytes.NewBuffer(postBody))
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

	collect := Objects.ResponseCollect{}

	err = json.Unmarshal(body, &collect)
	if err != nil {
		fmt.Println(err)
	}

	return collect
}

func PlayerReset() Objects.ResponseReset {
	client := http.Client{}
	request, err := http.NewRequest("DELETE", URL+"/player/reset", bytes.NewBuffer(nil))
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

	reset := Objects.ResponseReset{}

	err = json.Unmarshal(body, &reset)
	if err != nil {
		fmt.Println(err)
	}

	return reset
}

func PlayerRound() Objects.ResponseRound {
	client := http.Client{}
	request, err := http.NewRequest("GET", URL+"/player/rounds", bytes.NewBuffer(nil))
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

	round := Objects.ResponseRound{}

	err = json.Unmarshal(body, &round)
	if err != nil {
		fmt.Println(err)
	}

	return round
}
