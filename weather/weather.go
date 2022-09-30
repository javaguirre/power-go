package weather

import (
	"encoding/json"
	"fmt"
)

type Conditions struct {
	Summary string
}

type OWMResponse struct {
	Weather []struct {
		Main string
	}
}

func ParseResponse(data []byte) (Conditions, error) {
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

func FormatURL(location string, key string) string {
	baseURL := "https://api.openweathermap.org/data/2.5/weather"
	return fmt.Sprintf("%s?q=%s&appid=%s", baseURL, location, key)
}
