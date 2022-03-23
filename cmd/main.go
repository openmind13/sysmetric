package main

import (
	"log"
	"time"

	"sysmetric/sysmonitor"
	"sysmetric/sysmonitor/resources/network"
)

const version = "v0.0.0-in-docker"

// func main() {
// 	log.Println("Sysmonitor")

// 	systemMonitor, err := sysmonitor.New(sysmonitor.Config{
// 		ScrapePeriod: time.Second,
// 	})
// 	go systemMonitor.Start()

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

func main() {
	log.Println("sysmonitor: ", version)

	realIfc, _, err := network.GetNetworkInterfaces()
	if err != nil {
		log.Fatal(err)
	}

	systemMonitor, err := sysmonitor.New(sysmonitor.Config{
		ScrapePeriod: time.Second,
		NetworkConfig: network.Config{
			Period:       time.Second,
			NetInterface: realIfc[0],
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	go systemMonitor.Start()

	select {}
}
