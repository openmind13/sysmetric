package cpu

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var (
	lastUsage = time.Now()
)

func GetUsagePercent() (float64, error) {
	defer func() {
		lastUsage = time.Now()
	}()
	percent, err := cpu.Percent(time.Since(lastUsage), false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

type Monitor struct{}

func NewMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) GetCpuUsagePercent() (float64, error) {
	defer func() {
		lastUsage = time.Now()
	}()
	percent, err := cpu.Percent(time.Since(lastUsage), false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}
