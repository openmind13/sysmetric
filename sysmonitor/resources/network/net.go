package network

import (
	"log"
	"sync"
	"time"

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
	mu                sync.Mutex
	config            Config
	calcTime          time.Time
	networkInterfaces []networkInterface
	stat              Stat
}

func NewNetworkMonitor(config Config) (*NetworkMonitor, error) {
	m := &NetworkMonitor{
		config: config,
	}

	realIfcs, _, err := getNetworkInterfaces()
	if err != nil {
		return nil, err
	}
	m.networkInterfaces = realIfcs
	return m, nil
}

func (m *NetworkMonitor) Start() {
	lastTxBytesPerInterface := map[string]uint64{}
	lastRxBytesPerInterface := map[string]uint64{}

	deltaTxBytes := map[string]uint64{}
	deltaRxBytes := map[string]uint64{}

	period := time.Second
	netIfcs := m.networkInterfaces

	for {
		stat := Stat{
			RxBytesPerSeconds: map[string]float64{},
			TxBytesPerSeconds: map[string]float64{},
		}

		networkStat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
		if err != nil {
			log.Fatal(err)
		}
		for _, netStat := range networkStat {
			ifcName := netStat.Iface
			for _, ifc := range netIfcs {
				if ifcName == ifc.Name {
					deltaTxBytes[ifcName] = netStat.TxBytes - lastTxBytesPerInterface[ifcName]
					deltaRxBytes[ifcName] = netStat.RxBytes - lastRxBytesPerInterface[ifcName]

					lastTxBytesPerInterface[ifcName] = netStat.TxBytes
					lastRxBytesPerInterface[ifcName] = netStat.RxBytes

					stat.TxBytesPerSeconds[ifcName] = float64(deltaTxBytes[ifcName]/1000) / period.Seconds()
					stat.RxBytesPerSeconds[ifcName] = float64(deltaRxBytes[ifcName]/1000) / period.Seconds()
				}
			}
		}

		m.mu.Lock()
		m.stat = stat
		netIfcs = m.networkInterfaces
		period := m.config.Period
		m.mu.Unlock()

		time.Sleep(period)
	}
}

func (m *NetworkMonitor) GetStats() Stat {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.stat
}
