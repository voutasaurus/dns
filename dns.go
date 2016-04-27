package main

import (
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatal("Require hostname")
	}
	log.Println("IP addr:", lookup(os.Args[1]))
	log.Println("CName:", cname(os.Args[1]))
}

func lookup(host string) []string {
	addr, err := net.LookupHost(host)
	if err != nil {
		log.Fatal(err)
	}
	return addr
}

func cname(host string) string {
	cn, err := net.LookupCNAME(host)
	if err != nil {
		log.Fatal(err)
	}
	return cn
}
