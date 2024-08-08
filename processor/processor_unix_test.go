//go:build linux || darwin
// +build linux darwin

// SPDX-License-Identifier: MIT

package processor

import (
	"testing"
)

func TestScaleWorkersToLimit(t *testing.T) {
	scaleWorkersToLimit(10, 10)
}

func TestConfigureLimitsUnix(t *testing.T) {
	ConfigureLimitsUnix()
}

func TestConfigureLimitsUnixHighWater(t *testing.T) {
	FileProcessJobWorkers = 1
	DirectoryWalkerJobWorkers = 1
	ConfigureLimitsUnix()
}
