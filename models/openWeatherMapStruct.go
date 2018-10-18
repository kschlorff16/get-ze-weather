package models

import (
	"encoding/json"
	"fmt"
)

type openWeatherMapAPIResponse struct {
	City string     `json:"name"`
	Main mainStruct `json:"main"`
}

type mainStruct struct {
	Temperature float32 `json:"temp"`
}

func unmarshalObjectFromJSON(httpResponseJSON []byte) openWeatherMapAPIResponse {
	var response openWeatherMapAPIResponse

	err := json.Unmarshal(httpResponseJSON, &response)

	if err != nil {
		fmt.Println(err)
	}

	return response
}

/*
PrintLocationAndTemperatureInfo serves as a convenience accessor to the api response object. While not exceptionally useful here, encapsulating the struct inside of a separate package and giving it some useful exported and unexported functions means that we've set ourselves up for more flexibility down the road.
*/
func PrintLocationAndTemperatureInfo(httpResponseJSON []byte) {
	var finalObject = unmarshalObjectFromJSON(httpResponseJSON)
	fmt.Printf("%s weather:\n", finalObject.City)
	fmt.Printf("%v degrees Fahrenheit\n", finalObject.Main.Temperature)
}
