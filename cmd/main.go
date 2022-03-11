package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"sysmetric/procmonitor"
	"sysmetric/sysinformer"
	"sysmetric/sysmonitor"
	"time"
)

func main() {
	// go startSystemMonitor()
	// go startProcMonitor()
	go startSystemInformer()

	stopChan := make(chan struct{}, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stopChan:
		log.Println("Stopped")
		os.Exit(0)
	case sig := <-sigChan:
		fmt.Println("Signal:", sig)
		os.Exit(0)
	}
}

func startSystemMonitor() {
	systemMonitor := sysmonitor.New()
	log.Println("Starting system monitor")
	go systemMonitor.Start()

	for {
		stats := <-systemMonitor.StatisticChan
		// fmt.Printf("%+v\n\n", stats)
		for _, netStat := range stats.NetworkStat {
			fmt.Println(netStat.Iface, netStat.RxBytes)
		}

		fmt.Println()
		time.Sleep(time.Second)
	}
}

func startProcMonitor() {
	procMonitor := procmonitor.New()
	log.Println("Starting proc monitor")
	go procMonitor.Start()

	for {
		time.Sleep(time.Second)
	}
}

func startSystemInformer() {
	systemInformer := sysinformer.New()
	log.Println("Starting system informer")
	go systemInformer.Start()

	for {
		sysInfo := <-systemInformer.SystemInfoChan
		for _, nic := range sysInfo.Network.NICs {
			// fmt.Printf("%+v\n", nic.)
			// fmt.Println(nic.String())
			fmt.Println(nic.Name, " - ", nic.MacAddress, nic.IsVirtual)
			// for _, cap := range nic.Capabilities {
			// 	fmt.Print(cap.Name, " ", cap.IsEnabled, " ", cap.CanEnable, "\t")
			// }
		}
		// fmt.Printf("%+v\n\n", sysInfo.Network.NICs, sysInfo.Network.)

		fmt.Println()
		time.Sleep(time.Second)
	}
}
