package main

import (
	"fmt"
	"log"
	"net"
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

	ifcs, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, ifc := range ifcs {
		fmt.Println(ifc.Name, ifc.Flags)
	}
}
