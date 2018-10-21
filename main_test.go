package main

import (
	"os"
	"strings"
	"testing"
)

func TestResponseIsEmpty(t *testing.T) {
	emptyInput := ""
	if !responseIsEmpty(emptyInput) {
		t.Errorf("The string was not empty. The string's value was %v", emptyInput)
	}
	notEmptyInput := "Hello"
	if responseIsEmpty(notEmptyInput) {
		t.Errorf("The string was empty. The string's value was %v", notEmptyInput)
	}
}

func TestResponseContainsOnlyLettersAndSpacesWithValidInputs(t *testing.T) {
	validResponses := []string{"london", "Chicago", "Des Moines", "Santa Monica"}

	for _, validInput := range validResponses {
		if !userInputContainsOnlyLettersAndSpaces(validInput) {
			t.Errorf("The input was not valid. The value that was invalid was %v.", validInput)
		}
	}
}

func TestUserInputContainsOnlyLettersAndSpacesWithInvalidInputs(t *testing.T) {
	invalidResponses := []string{"1l0v3c4ts"}

	for _, validInput := range invalidResponses {
		if userInputContainsOnlyLettersAndSpaces(validInput) {
			t.Errorf("The response was not valid. The value that was invalid was %v.", validInput)
		}
	}
}

func TestBuildFullURL(t *testing.T) {
	input := "Testing"
	os.Setenv("OPENWEATHERMAPAPIKEY", "definitelyanapikey")
	const expected = "http://api.openweathermap.org/data/2.5/weather?q=Testing&units=imperial&APPID=definitelyanapikey"
	actual := buildFullURL(input)
	if actual != expected {
		t.Errorf("The API URL was not built as expected.\n Expected: %v\n Actual: %v", expected, actual)
	}
}

func TestSplitStringOnComma(t *testing.T) {
	const beep = "beep,boop"
	if splitStringOnComma(beep) != "beep" {
		t.Errorf("The string was not split correctly. The expected value was beep, but got %v", splitStringOnComma(beep))
	}
}

func TestSplitStringOnSpace(t *testing.T) {
	const beep = "beep boop"
	if splitStringOnSpace(beep) != "beep" {
		t.Errorf("The string was not split correctly. The expected value was beep, but got %v", splitStringOnSpace(beep))
	}
}

func TestRemoveCommasAndSpaces(t *testing.T) {
	userResponses := []string{"london,uk", "Chicago", "Charleston", "ames, iowa"}

	for _, response := range userResponses {
		cleanedResponse := removeCommasAndSpaces(response)
		if strings.Contains(cleanedResponse, ",") || strings.Contains(cleanedResponse, " ") {
			t.Errorf("The input was not valid. The value that was invalid was %v.", cleanedResponse)
		}
	}
}

func TestProcessCitiesWithSpacesInTheNames(t *testing.T) {
	inputs := []string{"Des Moines", "Des Moines, IA", "San Juan", "Des Moines,    Ia", "Des Moines Ia US World"}

	for _, input := range inputs {
		result := processCitiesWithSpacesInTheNames(input)
		if strings.Count(result, " ") > 1 {
			t.Errorf("There are more than one spaces in the result. The input that produced this failing test was: %v", input)
		}
	}
}

func TestCityHasPrefix(t *testing.T) {
	inputs := []string{"San Jose", "Des Moines", "San Juan", "Santa Monica", "Cedar Rapids", "Council Bluffs"}

	for _, input := range inputs {
		cityHasPrefix := cityHasPrefix(input)
		if !cityHasPrefix {
			t.Errorf("There is a prefix in the city name. The input that produced this failing tests was: %v", input)
		}
	}
}
