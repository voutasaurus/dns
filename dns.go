package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()
	host := flag.Arg(0)
	if host == "" {
		fmt.Fprintln(os.Stderr, "usage: dns hostname")
		return
	}

	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprint(os.Stderr, "IP addrs: ")
	fmt.Println(addrs)

	cname, err := net.LookupCNAME(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprint(os.Stderr, "CNAME: ")
	fmt.Println(cname)
}
