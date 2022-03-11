package sysinformer

import (
	"log"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/bios"
	"github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jaypipes/ghw/pkg/memory"
	"github.com/jaypipes/ghw/pkg/net"
)

type SystemInfo struct {
	Cpu     *cpu.Info
	Memory  *memory.Info
	Network *net.Info
	Bios    *bios.Info
}

type SystemInformer struct {
	SystemInfoChan chan SystemInfo
}

func New() *SystemInformer {
	informer := &SystemInformer{
		SystemInfoChan: make(chan SystemInfo, 1),
	}
	return informer
}

func (i *SystemInformer) Start() {
	for {
		systemInfo := SystemInfo{}

		memory, err := ghw.Memory()
		if err != nil {
			log.Println(err)
		}
		systemInfo.Memory = memory

		cpu, err := ghw.CPU()
		if err != nil {
			log.Println(err)
		}
		systemInfo.Cpu = cpu

		bios, err := ghw.BIOS()
		if err != nil {
			log.Println(err)
		}
		systemInfo.Bios = bios

		network, err := ghw.Network()
		if err != nil {
			log.Println(err)
		}
		systemInfo.Network = network

		i.SystemInfoChan <- systemInfo
	}
}
