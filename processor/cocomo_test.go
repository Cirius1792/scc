// SPDX-License-Identifier: MIT

package processor

import (
	"testing"
)

func TestEstimateCost(t *testing.T) {
	eff := EstimateEffort(26, 1)
	got := EstimateCost(eff, 56000, 2.4)

	// Should be around 582
	if got < 580 || got > 585 {
		t.Errorf("Got %f", got)
	}
}

func TestEstimateCostManyLines(t *testing.T) {
	eff := EstimateEffort(77873, 1)
	got := EstimateCost(eff, 56000, 2.4)

	// Should be around 2602096
	if got < 2602000 || got > 2602100 {
		t.Errorf("Got %f", got)
	}
}

func TestEstimateScheduleMonths(t *testing.T) {
	eff := EstimateEffort(537, 1)
	got := EstimateScheduleMonths(eff)

	// Should be around 2.7
	if got < 2.6 || got > 2.8 {
		t.Errorf("Got %f", got)
	}
}
