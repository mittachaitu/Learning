package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/mittachaitu/Learning/grpc-k8s/pkg/server"
)

func main() {
	fmt.Println("Starting RPC service to execute comamnds")
	var wg sync.WaitGroup
	address, err := getAddress()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	server := server.NewServer(address, "35220")
	wg.Add(1)
	go func() {
		server.Start()
		wg.Done()
	}()
	wg.Wait()
}

func getAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("Failed to get ifaces error: %v", err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", fmt.Errorf("Failed to get address error: %v", err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", fmt.Errorf("Failed to get IPAddress")
}
