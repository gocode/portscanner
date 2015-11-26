package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var (
	host     string
	from, to int
	timeout  int
)

func init() {
	flag.StringVar(&host, "host", "localhost", "host name of the machine to scan")
	flag.IntVar(&from, "from", 0, "port from which to scan")
	flag.IntVar(&to, "to", 0, "port to which to scan")
	flag.IntVar(&timeout, "timeout", 5, "duration in seconds after which the connection should timeout")
}

func main() {
	flag.Parse()

	for i := from; i <= to; i++ {
		url := net.JoinHostPort(host, fmt.Sprint(i))
		_, err := net.DialTimeout("tcp", url, time.Duration(timeout)*time.Second)
		if err == nil {
			fmt.Println(i)
		}
	}
}
