package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

type status uint

const (
	line = "+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+=+"
	statusTimeout status = iota
	statusOpen
)

var (
	host string
	min int
	max int
	timeout int
	suppress bool
)

func init() {
	flag.StringVar(&host, "host", "", "The host address")
	flag.IntVar(&min, "min", 0, "The minimum port number (inclusive)")
	flag.IntVar(&max, "max", 0, "The maximum port number (inclusive)")
	flag.IntVar(&timeout, "timeout", 1000, "The timeout in milliseconds")
	flag.BoolVar(&suppress, "suppress", false, "Whether to suppress verbose output")
	flag.Parse()
}

func main() {
	if host == "" {
		fmt.Println("Illegal host name.")
		return
	}
	if !isHostValid() {
		fmt.Println("Unable to resolve host name.")
		return
	}
	if min == 0 || max == 0 {
		fmt.Printf("Illegal port number(s): min=%d max=%d.\n", min, max)
		return
	}
	if min > max {
		fmt.Println("Illegal port range.")
		return
	}
	total := max - min + 1
	suffix := ""
	if total > 1 {
		suffix = "s"
	}
	fmt.Println(line)
	fmt.Printf("Scanning %d port%s @ %s.\n", total, suffix, host)
	fmt.Printf("- Minimum port is %d.\n", min)
	fmt.Printf("- Maximum port is %d.\n", max)
	fmt.Println(line)
	fmt.Println("Working...")
	open := make([]int, 0)
	for port := min; port <= max; port++ {
		var message string
		switch getStatus(port) {
		case statusTimeout:
			message = "closed"
		case statusOpen:
			message = "open"
			open = append(open, port)
		default:
			message = "error"
		}
		if !suppress {
			fmt.Printf("The port %d is %s.\n", port, message)
		}
	}
	fmt.Println(line)
	fmt.Printf("Total open ports: %d.\n", len(open))
	for _, port := range open {
		fmt.Printf("- %d\n", port)
	}
	fmt.Println(line)
}

func isHostValid() bool {
	_, err := net.ResolveIPAddr("", host)
	return err == nil
}

func getStatus(port int) status {
	con, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), time.Duration(timeout) * time.Millisecond)
	if err != nil {
		if opErr, success := err.(*net.OpError); !success || !opErr.Timeout() {
			fmt.Printf("Cannot handle error %T.\n", err)
		}
		return statusTimeout
	}
	if con != nil {
		con.Close()
	}
	return statusOpen
}