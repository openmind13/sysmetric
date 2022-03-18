package main

import (
	"fmt"
	"log"
	"time"

	"sysmetric/procmonitor"
	"sysmetric/sysinformer"
	"sysmetric/sysmonitor"
	"sysmetric/sysmonitor/netmonitor"
	"sysmetric/sysmonitor/netmonitor/ifc"
)

const version = "v0.0.0-in-docker"

func main() {
	realIfcs, virtualIfcs, err := ifc.GetInterfaces()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("real:", realIfcs)
	fmt.Println("virtual:", virtualIfcs)
}

// func main() {
// 	log.Println("Sysmonitor")
// 	// go startSystemMonitor()
// 	// go startProcMonitor()
// 	// go startSystemInformer()
// 	// go startNetMonitor()

// 	ifaces, err := netmonitor.GetInterfacesSpeed("enp2s0", "docker0")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%+v\n", ifaces)

// 	errChan := make(chan error, 1)
// 	stopChan := make(chan struct{}, 1)
// 	sigChan := make(chan os.Signal, 1)
// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

// 	httpServer := httpserver.New(httpserver.Config{
// 		HttpMetricAddr: "0.0.0.0:3001",
// 	})
// 	go httpServer.Start(errChan)

// 	metric.Info.WithLabelValues(version).Set(1)

// 	select {
// 	case err := <-errChan:
// 		log.Fatal(err)
// 	case <-stopChan:
// 		log.Println("Stopped")
// 		os.Exit(0)
// 	case sig := <-sigChan:
// 		fmt.Println("Signal:", sig)
// 		os.Exit(0)
// 	}
// }

func startSystemMonitor() {
	systemMonitor := sysmonitor.New()
	log.Println("Starting system monitor")
	go systemMonitor.Start()

	// for {
	// 	statistic := <-systemMonitor.StatisticChan
	// 	// fmt.Printf("%+v\n\n", stats)
	// 	for _, netStat := range statistic.NetworkStat {
	// 		fmt.Println(netStat.Iface, netStat.RxBytes)
	// 	}

	// 	fmt.Println()
	// 	time.Sleep(time.Second)
	// }

	select {}
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

func startNetMonitor() {
	netMonitor := netmonitor.NewNetworkMonitor(netmonitor.Config{
		Period: time.Second,
	})
	netMonitor.Run1()

	for {
		// stats := netMonitor.GetStats()
		// fmt.Printf("%+v\n", stats)

		netMonitor.GetStats()
		time.Sleep(2 * time.Second)
	}
}
