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

// FIXME: commented out until I can run it on Mac or switch the command
// under another OS
// func TestGetPmSetOutput(t *testing.T) {
// 	t.Parallel()
// 	text, err := battery.GetPmsetOutput()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	status, err := battery.ParsePmsetOutput(text)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	t.Logf("Charge: %d%%", status.ChargePercent)
// }
