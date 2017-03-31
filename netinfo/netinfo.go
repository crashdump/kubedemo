package netinfo

import (
	"os"
	"net"
)

type Network struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
}

func GetNetwork() Network {
	network := Network{}

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	network.Hostname = hostname

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	var addr string = ""
	for _, addrs := range addrs {
		addr = addrs.String()
	}

	network.IP = addr

	return network
}

