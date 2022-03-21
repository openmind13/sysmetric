package sysmonitor

import (
	"fmt"
	"sysmetric/sysmonitor/resources/cpu"
	"sysmetric/sysmonitor/resources/memory"
	"time"
)

type Statistic struct {
	CpuUsagePercent     float64
	MemoryUsagePercent  float64
	NetworkUsagePercent float64
}

type Config struct {
	ScrapePeriod            time.Duration
	CpuCollectingPeriod     time.Duration
	MemoryCollectingPeriod  time.Duration
	NetworkCollectingPeriod time.Duration
}

type SystemMonitor struct {
	config     Config
	configChan chan Config

	cpuMonitor *cpu.Monitor
	memMonitor *memory.Monitor
}

func New(config Config) *SystemMonitor {
	systemMonitor := SystemMonitor{
		config:     config,
		configChan: make(chan Config, 1),
		cpuMonitor: cpu.NewMonitor(),
		memMonitor: memory.NewMemoryMonitor(),
	}
	return &systemMonitor
}

func (m *SystemMonitor) Start() {
	for {
		// select {
		// case config := <-m.configChan:
		// 	m.config = config
		// }

		cpuUsage, _ := m.cpuMonitor.GetCpuUsagePercent()
		memUsage, _ := m.memMonitor.GetMemoryUsagePercent()
		fmt.Println(cpuUsage, memUsage)

		time.Sleep(m.config.ScrapePeriod)
	}
}

func (m *SystemMonitor) SetConfig(config Config) {
	m.configChan <- config
}
