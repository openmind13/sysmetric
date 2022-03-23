package network

import (
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type Config struct {
	NetInterface NetInterface
	Period       time.Duration
}

var errTemplate = "error in network monitor config. field: %s"

func (c *Config) Validate() error {
	if c.Period == 0 {
		return fmt.Errorf(errTemplate, "period")
	}
	return nil
}

type Stat struct {
	RxUsagePercent float64
	TxUsagePercent float64
}

type Monitor struct {
	mu     sync.Mutex
	config Config
	netIfc NetInterface
	stat   Stat
}

func NewNetworkMonitor(config Config) (*Monitor, error) {
	m := &Monitor{
		config: config,
		netIfc: config.NetInterface,
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	netIfc := m.netIfc

	var lastTxBytes, lastRxBytes uint64 = 0, 0
	var deltaTxBytes, deltaRxBytes uint64 = 0, 0
	var deltaTxBits, deltaRxBits uint64 = 0, 0
	var currTxSpeedBitsps, currRxSpeedBitsps float64 = 0.0, 0.0

	networkStat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
	if err != nil {
		return nil, err
	}

	for _, netStat := range networkStat {
		if netStat.Iface == netStat.Iface {
			lastRxBytes = netStat.RxBytes
			lastTxBytes = netStat.TxBytes
		}
	}

	period := m.config.Period
	time.Sleep(period)

	go func() {
		for {
			stat := Stat{}

			networkStat, err := linuxproc.ReadNetworkStat("/proc/net/dev")
			if err != nil {
				log.Fatal("Failed to read /proc/net/dev", err)
			}

			for _, netStat := range networkStat {
				ifcName := netStat.Iface
				if ifcName == netIfc.Name {
					deltaTxBytes = netStat.TxBytes - lastTxBytes
					deltaRxBytes = netStat.RxBytes - lastRxBytes

					// fmt.Println("deltaTxBytes:", deltaTxBytes, "deltaRxBytes:", deltaRxBytes)
					// fmt.Println("deltaTxKBytes:", deltaTxBytes/1000.0, "deltaRxKBytes:", deltaRxBytes/1000.0)
					// fmt.Println("deltaTxBBytes:", deltaTxBytes/1000.0/1000.0, "deltaRxMBytes:", deltaRxBytes/1000.0/1000.0)

					lastTxBytes = netStat.TxBytes
					lastRxBytes = netStat.RxBytes

					deltaTxBits = deltaTxBytes * 8.0
					deltaRxBits = deltaRxBytes * 8.0

					// fmt.Println("deltaTxMBits:", deltaTxBits/1000/1000, "deltaRxMBits:", deltaRxBits/1000/1000)

					currTxSpeedBitsps = float64(deltaTxBits) / period.Seconds()
					currRxSpeedBitsps = float64(deltaRxBits) / period.Seconds()

					// fmt.Println("txSpeed:", currTxSpeedBitsps/1000/1000, "rxSpeed:", currRxSpeedBitsps/1000/1000)

					txUsagePercent := currTxSpeedBitsps / float64(netIfc.BandwidthMbs*1000)
					rxUsagePercent := currRxSpeedBitsps / float64(netIfc.BandwidthMbs*1000)

					stat.TxUsagePercent = math.Round(txUsagePercent*1000) / 1000
					stat.RxUsagePercent = math.Round(rxUsagePercent*1000) / 1000
				}
			}

			m.mu.Lock()
			m.stat = stat
			netIfc = m.netIfc
			period := m.config.Period
			m.mu.Unlock()

			time.Sleep(period)
		}
	}()

	return m, nil
}

func (m *Monitor) GetStats() Stat {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.stat
}
