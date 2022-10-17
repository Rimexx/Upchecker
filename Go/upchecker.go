package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"sync"
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

	var wg sync.WaitGroup
	wg.Add(len(targets))

	for _, target := range targets {
		go func(target string) {
			defer wg.Done()
			splitted := strings.Split(target, ":")
			host := splitted[0]
			port := splitted[1]
			open := probeAddress("tcp", target)
			status := "FAIL"
			if open {
				status = "OK"
			}
			fmt.Printf("Probing target %s -- Host %s on TCP port %s ... %s\n", target, host, port, status)

		}(target)

	}
	wg.Wait()
}
