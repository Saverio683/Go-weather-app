package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RespBody struct { //struct fields must match the api ones
	Weather   []any
	Main, Sys interface{}
	Dt        int
}

func parseJson(jsonString io.Reader) (RespBody, error) { //parse the json response
	body := RespBody{} //the request body

	err := json.NewDecoder(jsonString).Decode(&body) //if the function succeeds, it will save the result in the RespBody struct{
	if err != nil {
		log.Println("error decoding response body", err)
		return RespBody{}, err
	}

	return body, nil
}

const id = "17e57c3ab6fdc68fb851ae80c2f9c4b6"

func MakeRequest(city, nation string) (RespBody, error) { //makes the request to the OpenWeatherMap API
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "," + nation + "&units=metric&appid=" + id
	resp, err := http.Get(url)

	if err != nil {
		log.Println("error retrieving weather api response", err)
		return RespBody{}, err
	}
	defer resp.Body.Close()

	parsedResp, err := parseJson(resp.Body)
	if err != nil {
		log.Println("error decoding response body")
		return RespBody{}, err
	}

	return parsedResp, nil
}
