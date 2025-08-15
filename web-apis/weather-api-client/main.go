package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WeatherData struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
		Humidity int `json:"humidity"`
		WindKph float64 `json:"wind_kph"`
	} `json:"current"`
}

func main() {
	// Check if city name is provided
	if len(os.Args) < 2 {
		fmt.Println("Please provide a city name as an argument")
		fmt.Println("Usage: go run . <city>")
		os.Exit(1)
	}

	city := os.Args[1]
	
	// Note: This is a demo API key that may not work
	// In a real application, you would need to get your own API key from https://www.weatherapi.com/
	apiKey := "YOUR_API_KEY_HERE"
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", apiKey, city)

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		os.Exit(1)
	}

	// Check if response is successful
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", string(body))
		os.Exit(1)
	}

	// Parse JSON response
	var weather WeatherData
	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	// Display weather information
	fmt.Printf("Weather in %s, %s:\n", weather.Location.Name, weather.Location.Country)
	fmt.Printf("Temperature: %.1f°C\n", weather.Current.TempC)
	fmt.Printf("Condition: %s\n", weather.Current.Condition.Text)
	fmt.Printf("Humidity: %d%%\n", weather.Current.Humidity)
	fmt.Printf("Wind: %.1f km/h\n", weather.Current.WindKph)
}
