package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

func main() {
	// netInfo, err := ghw.Network()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, ifc := range netInfo.NICs {
	// 	fmt.Println(ifc.Name, ifc.IsVirtual, ifc.MacAddress)
	// 	for _, cap := range ifc.Capabilities {
	// 		fmt.Println(cap.Name, cap.IsEnabled)
	// 	}
	// }

	// realIfcs, virtualIfcs, err := ifc.GetInterfaces()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("real:", realIfcs)
	// fmt.Println("virtual:", virtualIfcs)

	// ifcs, err := net.Interfaces()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, ifc := range ifcs {
	// 	fmt.Println(ifc.Name, ifc.Flags)
	// }

	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}
