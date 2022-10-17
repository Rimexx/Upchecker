package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func probePort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func main() {
	fmt.Println("Port Scanning")

	host := "localhost"
	port := "1337"
	name := host + port

	open := probePort("tcp", "localhost", 1337)
	fmt.Printf("Probing target %s -- Host %s on TCP port %s ... ", host, port, name)
	fmt.Printf("Port Open: %t\n", open)
}
