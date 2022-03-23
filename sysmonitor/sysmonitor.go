package sysmonitor

import (
	"errors"
	"fmt"
	"time"

	"sysmetric/sysmonitor/resources/cpu"
	"sysmetric/sysmonitor/resources/memory"
	"sysmetric/sysmonitor/resources/network"
)

type Statistic struct {
	CpuUsagePercent     float64
	MemoryUsagePercent  float64
	NetworkUsagePercent float64
}

type Config struct {
	ScrapePeriod  time.Duration
	NetworkConfig network.Config
}

var errTemplate = "error in system monitor config: field: %s"

func (c *Config) Validate() error {
	if c.ScrapePeriod == 0 {
		return fmt.Errorf(errTemplate, "scrape_period")
	}
	return nil
}

type SystemMonitor struct {
	config     Config
	configChan chan Config

	cpuMonitor *cpu.Monitor
	memMonitor *memory.Monitor
	netMonitor *network.Monitor
}

func New(config Config) (*SystemMonitor, error) {
	systemMonitor := SystemMonitor{
		config:     config,
		configChan: make(chan Config, 1),

		cpuMonitor: cpu.NewMonitor(),
		memMonitor: memory.NewMemoryMonitor(),
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	realIfc, _, err := network.GetNetworkInterfaces()
	if err != nil {
		return nil, err
	}

	if len(realIfc) > 1 || len(realIfc) < 1 {
		return nil, errors.New("one net interface needed")
	}

	netMonitor, err := network.NewNetworkMonitor(systemMonitor.config.NetworkConfig)
	if err != nil {
		return nil, err
	}
	systemMonitor.netMonitor = netMonitor

	return &systemMonitor, nil
}

func (m *SystemMonitor) Start() {
	for {
		// select {
		// case config := <-m.configChan:
		// 	m.config = config
		// }

		// cpuUsage, _ := m.cpuMonitor.GetCpuUsagePercent()
		// memUsage, _ := m.memMonitor.GetMemoryUsagePercent()
		netStat := m.netMonitor.GetStats()
		fmt.Printf("%+v\n", netStat)
		// fmt.Println(cpuUsage, memUsage)

		time.Sleep(m.config.ScrapePeriod)
	}
}

func (m *SystemMonitor) SetConfig(config Config) {
	m.configChan <- config
}
