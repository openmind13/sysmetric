package netmonitor

import (
	"log"
	"sync"
	"time"

	"sysmetric/metric"
	"sysmetric/sysmonitor/netmonitor/ifc"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Config struct {
	Period time.Duration
}

type Stat struct {
	RxBytesPerSeconds map[string]float64
	TxBytesPerSeconds map[string]float64
}

type NetworkMonitor struct {
	mu                  sync.RWMutex
	config              Config
	calcTime            time.Time
	networkInterfaces   []ifc.Ifc
	rxBytesPerInterface map[string]uint64
	txBytesPerInterface map[string]uint64
	stat                Stat
}

func NewNetworkMonitor(config Config) *NetworkMonitor {
	m := &NetworkMonitor{
		config:              config,
		rxBytesPerInterface: map[string]uint64{},
		txBytesPerInterface: map[string]uint64{},
	}
	return m
}

func (m *NetworkMonitor) Run1() {
	calcTime := time.Now()
	networkstat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		log.Fatal(err)
	}

	m.mu.Lock()
	m.calcTime = calcTime
	for _, netstat := range networkstat {
		m.rxBytesPerInterface[netstat.Iface] = netstat.RxBytes
		m.txBytesPerInterface[netstat.Iface] = netstat.TxBytes
	}
	m.mu.Unlock()
}

func (m *NetworkMonitor) GetStats() Stat {
	stat := Stat{
		RxBytesPerSeconds: map[string]float64{},
		TxBytesPerSeconds: map[string]float64{},
	}

	getStatsCalcTime := time.Now()
	networkstat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		log.Fatal(err)
	}

	deltaTxBytes := map[string]uint64{}
	deltaRxBytes := map[string]uint64{}

	m.mu.RLock()
	for _, netstat := range networkstat {
		ifcName := netstat.Iface
		deltaRxBytes[ifcName] = netstat.RxBytes - m.rxBytesPerInterface[ifcName]
		deltaTxBytes[ifcName] = netstat.TxBytes - m.txBytesPerInterface[ifcName]

		stat.RxBytesPerSeconds[ifcName] = float64(deltaRxBytes[ifcName]) / float64(getStatsCalcTime.Sub(m.calcTime))
		stat.TxBytesPerSeconds[ifcName] = float64(deltaTxBytes[ifcName]) / float64(getStatsCalcTime.Sub(m.calcTime))

		metric.RxBytesPerSecond.WithLabelValues(ifcName).Set(stat.RxBytesPerSeconds[ifcName])
		metric.TxBytesPerSecond.WithLabelValues(ifcName).Set(stat.TxBytesPerSeconds[ifcName])
	}
	m.mu.RUnlock()

	return stat
}
