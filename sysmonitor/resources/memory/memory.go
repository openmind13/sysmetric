package memory

import "github.com/shirou/gopsutil/mem"

type Monitor struct{}

func NewMemoryMonitor() *Monitor {
	return &Monitor{}
}

func (m *Monitor) GetMemoryUsagePercent() (float64, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return v.UsedPercent, nil
}
