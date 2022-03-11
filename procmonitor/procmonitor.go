package procmonitor

import (
	"log"

	sigar "github.com/cloudfoundry/gosigar"
)

type Statistic struct {
	Data []ProcessStat
}

type ProcessStat struct {
	ProcMem   sigar.ProcMem
	ProcCpu   sigar.ProcCpu
	ProcArgs  sigar.ProcArgs
	ProcExe   sigar.ProcExe
	ProcState sigar.ProcState
	ProcTime  sigar.ProcTime
}

type ProcessMonitor struct {
	StatisticChan chan Statistic
}

func New() *ProcessMonitor {
	osmonitor := ProcessMonitor{
		StatisticChan: make(chan Statistic, 1),
	}
	return &osmonitor
}

func (m *ProcessMonitor) Start() {
	for {
		stats := Statistic{}

		pidList := []int{}

		for _, pid := range pidList {
			procMem := sigar.ProcMem{}
			procCpu := sigar.ProcCpu{}
			procArgs := sigar.ProcArgs{}
			procExe := sigar.ProcExe{}
			procState := sigar.ProcState{}
			procTime := sigar.ProcTime{}

			if err := procMem.Get(pid); err != nil {
				log.Println(err)
			}
			if err := procCpu.Get(pid); err != nil {
				log.Println(err)
			}
			if err := procArgs.Get(pid); err != nil {
				log.Println(err)
			}
			if err := procExe.Get(pid); err != nil {
				log.Println(err)
			}
			if err := procState.Get(pid); err != nil {
				log.Println(err)
			}
			if err := procTime.Get(pid); err != nil {
				log.Println(err)
			}

			processStat := ProcessStat{
				ProcMem:   procMem,
				ProcCpu:   procCpu,
				ProcArgs:  procArgs,
				ProcExe:   procExe,
				ProcState: procState,
				ProcTime:  procTime,
			}

			stats.Data = append(stats.Data, processStat)
		}

		m.StatisticChan <- stats
	}
}
