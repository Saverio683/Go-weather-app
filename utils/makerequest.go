package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	weatherdetail "weather-app/components/weather-detail"
)

type respBody struct { //struct fields must match the api ones
	Weather   [1]map[string]any
	Main, Sys map[string]any
}

func parseJson(jsonString io.Reader) (respBody, error) { //parse the json response
	body := respBody{} //the request body

	err := json.NewDecoder(jsonString).Decode(&body) //if the function succeeds, it will save the apiResult in the respBody struct{
	if err != nil {
		log.Println("error decoding response body", err)
		return respBody{}, err
	}

	return body, nil
}

func getStringFromAny(Value any) string {
	var parsedValue string
	if float, ok := Value.(float64); ok {
		parsedValue = strconv.FormatFloat(float, 'f', 0, 64)
	}

	if str, ok := Value.(string); ok {
		parsedValue = str
	}

	return parsedValue
}

const id = "17e57c3ab6fdc68fb851ae80c2f9c4b6"

func MakeRequest(city, nation string) ([]*weatherdetail.WeatherDetail, error) { //makes the request to the OpenWeatherMap API
	apiResult := make([]*weatherdetail.WeatherDetail, 0)
	url := "https://api.openweathermap.org/data/2.5/weather?q=" + city + "," + nation + "&units=metric&appid=" + id
	resp, err := http.Get(url)

	if err != nil {
		log.Println("error retrieving weather api response", err)
		return apiResult, err
	}
	defer resp.Body.Close()

	parsedResp, err := parseJson(resp.Body)
	if err != nil {
		log.Println("error decoding response body")
		return apiResult, err
	}

	apiResult = append(apiResult,
		/* 		weatherdetail.NewWeatherDetail("iconId", getStringFromAny(parsedResp.Weather[0]["icon"]), 0),
		   		weatherdetail.NewWeatherDetail("temp", getStringFromAny(parsedResp.Main["temp"]), 1), */
		weatherdetail.NewWeatherDetail("temp min", getStringFromAny(parsedResp.Main["temp_min"]), 0),
		weatherdetail.NewWeatherDetail("temp max", getStringFromAny(parsedResp.Main["temp_max"]), 1),
		weatherdetail.NewWeatherDetail("humidity", getStringFromAny(parsedResp.Main["humidity"]), 2),
		weatherdetail.NewWeatherDetail("pressure", getStringFromAny(parsedResp.Main["pressure"]), 3),
		weatherdetail.NewWeatherDetail("sunrise", getStringFromAny(parsedResp.Sys["sunrise"]), 4),
		weatherdetail.NewWeatherDetail("sunset", getStringFromAny(parsedResp.Sys["sunset"]), 5),
	)

	return apiResult, nil
}
