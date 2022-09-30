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
