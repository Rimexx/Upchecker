package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"
)

func probeAddress(protocol string, address string) bool {
	conn, err := net.DialTimeout(protocol, address, 2*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

type Targets []string

func (t *Targets) String() string {
	return fmt.Sprintln(*t)
}

func (t *Targets) Set(s string) error {
	*t = strings.Split(s, ",")
	return nil
}

func main() {
	var targets Targets
	flag.Var(&targets, "t", "a single host in the form: 127.0.0.1:8000, or a comma seperated list of hosts")
	flag.Parse()

	// TODO use go routines

	for _, target := range targets {
		splitted := strings.Split(target, ":")
		host := splitted[0]
		port := splitted[1]
		fmt.Printf("Probing target %s -- Host %s on TCP port %s ... ", target, host, port)
		open := probeAddress("tcp", target)
		fmt.Printf("Port Open: %t\n", open)
	}

}
