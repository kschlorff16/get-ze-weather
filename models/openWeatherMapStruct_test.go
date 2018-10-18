package models

import (
	"testing"
)

const sampleJSON = `{
    "coord": {
        "lon": -0.13,
        "lat": 51.51
    },
    "weather": [
        {
            "id": 804,
            "main": "Clouds",
            "description": "overcast clouds",
            "icon": "04n"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 54.99,
        "pressure": 1024,
        "humidity": 87,
        "temp_min": 51.8,
        "temp_max": 57.2
    },
    "visibility": 10000,
    "wind": {
        "speed": 3.36
    },
    "clouds": {
        "all": 92
    },
    "dt": 1539814800,
    "sys": {
        "type": 1,
        "id": 5091,
        "message": 0.0053,
        "country": "GB",
        "sunrise": 1539757755,
        "sunset": 1539795674
    },
    "id": 2643743,
    "name": "London",
    "cod": 200
}`

func TestUnMarshalObjectFromJSON(t *testing.T) {
	jsonAsByteArray := []byte(sampleJSON)
	actual := unmarshalObjectFromJSON(jsonAsByteArray)
	if actual.City != "London" {
		t.Errorf("The city was incorrect. Actual value was %v, expected value was London.", actual.City)
	}
	if actual.Main.Temperature != 54.99 {
		t.Errorf("The temperature was incorrect. Actual value was %v, expected value was 54.99.", actual.Main.Temperature)
	}
}
