package battery_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/power-go/battery"
)

func TestParsePmsetOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/pmset.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := battery.Status{
		ChargePercent: 100,
	}

	got, _ := battery.ParsePmsetOutput(string(data))

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
