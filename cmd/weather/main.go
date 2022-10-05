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
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s LOCATION\n\nExample %[1]s Cordoba, ES", os.Args[0])
	}

	location := os.Args[1]

	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the OPENWEATHERMAP_API_KEY!")
	}
	URL := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", BaseURL, location, key)

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
