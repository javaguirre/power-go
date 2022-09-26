package battery

import (
	"fmt"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

var pmSetOutput = regexp.MustCompile("([0-9]+)%")

func ParsePmsetOutput(output string) (Status, error) {
	matches := pmSetOutput.FindStringSubmatch(output)

	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse pmset, output: %q", output)
	}

	batteryPercent, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse chargepercentage: %q", matches[1])
	}

	return Status{
		ChargePercent: batteryPercent,
	}, nil
}
