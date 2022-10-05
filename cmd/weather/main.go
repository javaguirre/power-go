package main

import (
	"github.com/power-go/weather"
)

const BaseURL = "https://api.openweathermap.org"

func main() {
	weather.RunCLI()
}
