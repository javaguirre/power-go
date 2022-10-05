package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Conditions struct {
	Summary string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

type WeatherClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

func RunCLI() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s LOCATION\n\nExample %[1]s Cordoba, ES", os.Args[0])
	}

	location := os.Args[1]

	key := os.Getenv("OPENWEATHERMAP_API_KEY")
	if key == "" {
		log.Fatal("Please set the OPENWEATHERMAP_API_KEY!")
	}

	conditions, err := Get(location, key)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Conditions: %s\n", conditions.Summary)
}

func Get(location string, key string) (Conditions, error) {
	client := NewClient(key)
	conditions, err := client.GetWeather(location)

	if err != nil {
		return Conditions{}, err
	}

	return conditions, nil
}

func NewClient(key string) *WeatherClient {
	return &WeatherClient{
		APIKey:  key,
		BaseURL: "https://api.openweathermap.org",
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (w *WeatherClient) GetWeather(location string) (Conditions, error) {
	URL := w.FormatURL(location)
	data, err := w.MakeAPIRequest(URL)
	if err != nil {
		return Conditions{}, err
	}

	conditions, err := w.ParseResponse(data)
	if err != nil {
		return Conditions{}, err
	}

	return conditions, nil
}

func (w *WeatherClient) ParseResponse(data []byte) (Conditions, error) {
	var resp OWMResponse
	err := json.Unmarshal(data, &resp)
	if err != nil {
		return Conditions{}, fmt.Errorf(
			"invalid API response %s: %w", data, err)
	}

	if len(resp.Weather) < 1 {
		return Conditions{}, fmt.Errorf(
			"invalid API response %s: want at least one Weather element", data)
	}

	conditions := Conditions{
		Summary: resp.Weather[0].Main,
	}

	return conditions, nil
}

func (w *WeatherClient) FormatURL(location string) string {
	return fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s", w.BaseURL, location, w.APIKey)
}

func (w *WeatherClient) MakeAPIRequest(URL string) ([]byte, error) {
	resp, err := w.HTTPClient.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected Response status", resp.Status)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error when reading all data on Body")
		return nil, err
	}

	return data, nil
}
