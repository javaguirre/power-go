package weather_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/power-go/weather"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Conditions{
		Summary: "Clear",
	}

	got, err := weather.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponseEmpty(t *testing.T) {
	t.Parallel()

	_, err := weather.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestParseResponseInvalid(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}

	_, err = weather.ParseResponse(data)
	if err == nil {
		t.Fatal("want error parsing invalid response, got nil")
	}
}

func TestFormatURL(t *testing.T) {
	t.Parallel()
	location := "Cordoba,ES"
	key := "dummyAPIKey"
	want := "https://api.openweathermap.org/data/2.5/weather?q=Cordoba,ES&appid=dummyAPIKey"

	got := weather.FormatURL(location, key)

	if want != got {
		t.Fatalf("%s not equal, expected %s", got, want)
	}
}
