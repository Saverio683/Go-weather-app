package utils

import (
	"encoding/json"
	"io"
	"log"
	"strconv"
)

type respBody struct { //struct fields must match the api ones
	Weather   [1]map[string]any
	Main, Sys map[string]any
}

func ParseJson(jsonString io.Reader) (respBody, error) { //parse the json response
	body := respBody{} //the request body

	err := json.NewDecoder(jsonString).Decode(&body) //if the function succeeds, it will save the apiResult in the respBody struct{
	if err != nil {
		log.Println("error decoding response body", err)
		return respBody{}, err
	}

	return body, nil
}

func GetStringFromAny(Value any) string {
	var parsedValue string
	if float, ok := Value.(float64); ok {
		parsedValue = strconv.FormatFloat(float, 'f', 0, 64)
	}

	if str, ok := Value.(string); ok {
		parsedValue = str
	}

	return parsedValue
}
