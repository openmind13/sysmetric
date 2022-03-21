package network

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/net"
)

const (
	ETHTOOL_NAME = "ethtool"
)

type networkInterface struct {
	Name         string
	MacAddr      string
	BandwidthMbs int
	IsVirtual    bool
	Capabilities []*net.NICCapability
}

func getNetworkInterfaces() (realIfcs []networkInterface, virtualIfcs []networkInterface, err error) {
	netInfo, err := ghw.Network()
	if err != nil {
		return nil, nil, err
	}
	for _, nic := range netInfo.NICs {
		ifc := networkInterface{
			Name:         nic.Name,
			MacAddr:      nic.MacAddress,
			Capabilities: nic.Capabilities,
		}
		if nic.IsVirtual {
			ifc.IsVirtual = true
			ifc.BandwidthMbs = 0
			virtualIfcs = append(virtualIfcs, ifc)
			continue
		} else {
			ifc.IsVirtual = false
			bandwidth, err := getInterfaceBandwidth(ifc.Name)
			if err != nil {
				log.Fatal(err)
			}
			ifc.BandwidthMbs = bandwidth
			realIfcs = append(realIfcs, ifc)
		}
	}

	return realIfcs, virtualIfcs, nil
}

func getInterfaceBandwidth(ifcName string) (int, error) {
	out, err := exec.Command(ETHTOOL_NAME, ifcName).Output()
	if err != nil {
		return 0, err
	}
	buffer := bytes.Buffer{}
	buffer.Write(out)
	scanner := bufio.NewScanner(&buffer)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if !strings.Contains(scanner.Text(), "Speed") {
			continue
		}
		speedStrSplitted := strings.Split(scanner.Text(), " ")
		speedStr := strings.Split(speedStrSplitted[1], "M")[0]
		bandwidth, err := strconv.Atoi(speedStr)
		if err != nil {
			log.Println("Failed to parse interface bandwidth", ifcName)
			continue
		}
		return bandwidth, nil
	}
	return 0, errors.New("interface bandwidth unknown")
}
