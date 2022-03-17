package netmonitor

import (
	"log"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Stat struct{}

type Config struct {
	Period time.Duration
}

type NetworkMonitor struct {
	config            Config
	currentStat       Stat
	hookChan          chan struct{}
	incomingStatChan  chan Stat
	outcomingStatChan chan Stat
}

func NewNetworkMonitor(config Config) *NetworkMonitor {
	monitor := &NetworkMonitor{
		hookChan:          make(chan struct{}, 1),
		incomingStatChan:  make(chan Stat, 1),
		outcomingStatChan: make(chan Stat, 1),
	}

	go func() {
		for {
			select {
			case <-monitor.hookChan:
				monitor.outcomingStatChan <- monitor.currentStat
			case monitor.currentStat = <-monitor.incomingStatChan:
			}
		}
	}()

	go func() {
		for {
			networkStat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
			if err != nil {
				log.Fatal(err)
			}
			rxBytesPerInterface := map[string]uint64{}
			txBytesPerInterface := map[string]uint64{}

			for _, netstat := range networkStat {
				rxBytesPerInterface[netstat.Iface] = netstat.RxBytes
				txBytesPerInterface[netstat.Iface] = netstat.TxBytes
			}

			time.Sleep(monitor.config.Period)
		}
	}()

	return monitor
}

func (m *NetworkMonitor) GetStats() Stat {
	m.hookChan <- struct{}{}
	return <-m.outcomingStatChan
}
