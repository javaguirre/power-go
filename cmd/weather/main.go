package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/power-go/weather"
)

const BaseURL = "https://api.openweathermap.org"

func main() {
	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the OPENWEATHERMAP_API_KEY!")
	}
	URL := fmt.Sprintf("%s/data/2.5/weather?q=Cordoba,ES&appid=%s", BaseURL, key)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected Response status", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error when reading all data on Body")
	}

	conditions, err := weather.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Conditions: %s\n", conditions.Summary)
}
