package sysmonitor

import (
	"sysmetric/sysmonitor/netmonitor"
	"time"
)

// type Stats struct {
// 	Cpu  sigar.Cpu
// 	Mem  sigar.Mem
// 	Swap sigar.Swap
// }

// type Statistic struct {
// 	Stat        linuxproc.Stat
// 	Cpu         linuxproc.CPUStat
// 	CpuInfo     linuxproc.CPUInfo
// 	NetStat     linuxproc.NetStat
// 	NetworkStat []linuxproc.NetworkStat
// }

type SystemMonitor struct {
	// StatisticChan chan Statistic

	NetworkMonigor netmonitor.NetworkMonitor
}

func New() *SystemMonitor {
	systemMonitor := SystemMonitor{
		NetworkMonigor: *netmonitor.NewNetworkMonitor(netmonitor.Config{
			Period: time.Second,
		}),
	}
	return &systemMonitor
}

// func (m *SystemMonitor) Start() {
// 	for {
// 		mem := sigar.Mem{}
// 		cpu := sigar.Cpu{}
// 		swap := sigar.Swap{}
// 		cpuList := sigar.CpuList{}

// 		mem.Get()
// 		cpu.Get()
// 		cpuList.Get()
// 		swap.Get()

// 		stats := Stats{
// 			Cpu:  cpu,
// 			Mem:  mem,
// 			Swap: swap,
// 		}

// 		m.StatsChan <- stats

// 		fmt.Printf("%+v\n", stats)
// 	}
// }

func (m *SystemMonitor) Start() {

	for {
		m.NetworkMonigor.GetStats()
		// stat := m.NetworkMonigor.GetStats()
		// fmt.Printf("%+v\n", stat)
		time.Sleep(time.Second)
	}
	// for {
	// 	statistic := Statistic{}

	// 	stat, err := linuxproc.ReadStat("/proc/stat")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	cpuInfo, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// netStat, err := linuxproc.ReadNetStat("/proc/net/dev")
	// 	// if err != nil {
	// 	// 	log.Fatal(err)
	// 	// }

	// 	networkStat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	statistic.Stat = *stat
	// 	statistic.Cpu = stat.CPUStatAll
	// 	statistic.CpuInfo = *cpuInfo
	// 	// statistic.NetStat = *netStat
	// 	statistic.NetworkStat = networkStat

	// 	m.StatisticChan <- statistic
	// }
}
