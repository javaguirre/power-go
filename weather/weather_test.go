package weather_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/power-go/weather"
)

func TestParseResponse(t *testing.T) {
	t.Parallel()
	client := weather.NewClient("dummyAPIKey")
	data, err := os.ReadFile("testdata/weather.json")
	if err != nil {
		t.Fatal(err)
	}

	want := weather.Conditions{
		Summary: "Clear",
	}

	got, err := client.ParseResponse(data)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestParseResponseEmpty(t *testing.T) {
	t.Parallel()
	client := weather.NewClient("dummyAPIKey")

	_, err := client.ParseResponse([]byte{})
	if err == nil {
		t.Fatal("want error parsing empty response, got nil")
	}
}

func TestParseResponseInvalid(t *testing.T) {
	t.Parallel()
	client := weather.NewClient("dummyAPIKey")
	data, err := os.ReadFile("testdata/weather_invalid.json")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.ParseResponse(data)
	if err == nil {
		t.Fatal("want error parsing invalid response, got nil")
	}
}

func TestFormatURL(t *testing.T) {
	t.Parallel()
	client := weather.NewClient("dummyAPIKey")
	location := "Cordoba,ES"
	want := "https://api.openweathermap.org/data/2.5/weather?q=Cordoba,ES&appid=dummyAPIKey"

	got := client.FormatURL(location)

	if want != got {
		t.Fatalf("%s not equal, expected %s", got, want)
	}
}

func TestGetWeather(t *testing.T) {
	t.Parallel()

	ts := httptest.NewTLSServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open("testdata/weather.json")
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			io.Copy(w, f)
		}))
	defer ts.Close()

	client := weather.NewClient("dummyKey")
	client.BaseURL = ts.URL
	client.HTTPClient = ts.Client()

	location := "Mars"
	want := weather.Conditions{
		Summary: "Clear",
	}

	got, err := client.GetWeather(location)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
