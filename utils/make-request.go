package utils

import (
	"encoding/json"
	"io"
	"strconv"
)

type respBody struct { //struct fields must match the api ones
	Weather   [1]map[string]any //api values can be both numbers (float64) or strings, so that's why I used any
	Main, Sys map[string]any
}

func ParseJson(jsonString io.Reader) (respBody, error) { //parse the json response
	body := respBody{} //the request body

	err := json.NewDecoder(jsonString).Decode(&body) //if the function succeeds, it will save the apiResult in the respBody struct
	if err != nil {
		return respBody{}, err
	}

	return body, nil
}

func GetStringFromAny(value any) string { //converting to string
	var parsedValue string

	if float, ok := value.(float64); ok {
		parsedValue = strconv.FormatFloat(float, 'f', 0, 64)
	}
	if str, ok := value.(string); ok {
		parsedValue = str
	}

	return parsedValue
}
