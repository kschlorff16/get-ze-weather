package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/kschlorff16/get-ze-weather/models"
)

const baseURL = "http://api.openweathermap.org/data/2.5/weather?q="
const imperialUnits = "&units=imperial"
const appIDSuffix = "&APPID="
const prompt = "Where are you?"

func main() {
	if os.Getenv("OPENWEATHERMAPAPIKEY") == "" {
		fmt.Println("Whoa there. Make sure you set the OPENWEATHERMAPAPIKEY variable before you get started.")
		os.Exit(1)
	}
	fmt.Println(prompt)
	response := promptUserForInput()
	cleanedResponse := cleanUserResponse(response)
	if responseIsOk(cleanedResponse) {
		fullURL := buildFullURL(cleanedResponse)
		responseAsByteArray := convertHTTPResponseToByteArray(fullURL)
		models.PrintLocationAndTemperatureInfo(responseAsByteArray)
	}
}

func convertHTTPResponseToByteArray(fullURL string) []byte {
	httpResponse, err := http.Get(fullURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer httpResponse.Body.Close()
	httpResponseAsByteArray, _ := ioutil.ReadAll(httpResponse.Body)
	if strings.Contains(string(httpResponseAsByteArray), "city not found") {
		fmt.Println("That city wasn't found. Try again")
		os.Exit(1)
	}
	return httpResponseAsByteArray
}

func promptUserForInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if responseIsEmpty(text) {
		fmt.Println("You must enter a value.")
		os.Exit(1)
	}
	return strings.TrimSpace(text)
}

func cleanUserResponse(stringThatMayContainCommasOrSpaces string) string {
	cleanedInput := stringThatMayContainCommasOrSpaces
	if strings.Contains(stringThatMayContainCommasOrSpaces, ",") {
		cleanedInput = splitStringOnComma(stringThatMayContainCommasOrSpaces)
	}

	if strings.Contains(stringThatMayContainCommasOrSpaces, " ") {
		cleanedInput = splitStringOnSpace(stringThatMayContainCommasOrSpaces)
	}

	return cleanedInput
}

func responseIsOk(response string) bool {
	cleanedInput := response

	for _, character := range cleanedInput {
		if !unicode.IsLetter(character) {
			return false
		}
	}
	return true
}

func responseIsEmpty(response string) bool {
	if response == "" {
		return true
	}
	return false
}

func buildFullURL(location string) string {
	apiKey := os.Getenv("OPENWEATHERMAPAPIKEY")
	apiStrings := []string{baseURL, location, imperialUnits, appIDSuffix, apiKey}
	fmt.Println(strings.Join(apiStrings[:], ""))
	return strings.Join(apiStrings[:], "")
}

func splitStringOnComma(stringThatContainsCommas string) string {
	return strings.Split(stringThatContainsCommas, ",")[0]
}

func splitStringOnSpace(stringThatContainsSpaces string) string {
	fmt.Println("This is what I got: " + stringThatContainsSpaces)
	fmt.Printf("This is what I gave back: %v", strings.Split(stringThatContainsSpaces, " ")[0])
	return strings.Split(stringThatContainsSpaces, " ")[0]
}
