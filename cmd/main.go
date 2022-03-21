package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"sysmetric/httpserver"
	"sysmetric/metric"
	"sysmetric/sysmonitor"
)

const version = "v0.0.0-in-docker"

func main() {
	log.Println("Sysmonitor")

	systemMonitor := sysmonitor.New(sysmonitor.Config{
		ScrapePeriod: time.Second,
	})
	go systemMonitor.Start()

	errChan := make(chan error, 1)
	stopChan := make(chan struct{}, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	httpServer := httpserver.New(httpserver.Config{
		HttpMetricAddr: "0.0.0.0:3001",
	})
	go httpServer.Start(errChan)

	metric.Info.WithLabelValues(version).Set(1)

	select {
	case err := <-errChan:
		log.Fatal(err)
	case <-stopChan:
		log.Println("Stopped")
		os.Exit(0)
	case sig := <-sigChan:
		fmt.Println("Signal:", sig)
		os.Exit(0)
	}
}
