package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	protoCommand "github.com/mittachaitu/Learning/grpc-k8s/protobuffers/go/command"
	"google.golang.org/grpc"
	yaml "gopkg.in/yaml.v2"
)

var (
	ipAddress    string
	port         string
	commandsPath string
)

func init() {
	flag.StringVar(&ipAddress, "address", "", "gRPC server IP address")
	flag.StringVar(&port, "port", "35220", "gRPC server port")
	flag.StringVar(&commandsPath, "cmdpath", "/commands/config.yaml", "List of commands to query")
}

func main() {
	flag.Parse()
	conn, err := grpc.DialContext(context.TODO(), net.JoinHostPort(ipAddress, port), grpc.WithInsecure())
	if err != nil {
	}
	defer conn.Close()
	c := protoCommand.NewRunRPCCommandsClient(conn)
	requests, err := getCommandRequests(commandsPath)
	if err != nil {
		fmt.Printf("Failed to get command requests, error: %v\n", err)
		os.Exit(1)
	}
	// requests := []*protoCommand.CommandRequest{
	// 	{
	// 		Binary: "ls",
	// 		Args:   []string{"-lrtha", "/"},
	// 	},
	// 	{
	// 		Binary: "lscpu",
	// 		Args:   nil,
	// 	},
	// 	{
	// 		Binary: "nvme",
	// 		Args:   []string{"list", "-o json"},
	// 	},
	// 	{
	// 		Binary: "cat",
	// 		Args:   []string{"/proc/meminfo"},
	// 	},
	// }
	resp, err := c.RunShellCommands(context.TODO(), &protoCommand.ShellCommandsOutputRequest{Requests: requests})
	if err != nil {
		fmt.Println("Failed to execute gRPC call error: ", err)
		os.Exit(1)
	}
	for i := 0; i < len(requests); i++ {
		if resp.Responses[i].Output != nil {
			fmt.Printf("CMD: %s Args: %v Output: \n %s \n", requests[i].Binary, requests[i].Args, string(resp.Responses[i].Output))
		}
	}
	return
}

func getCommandRequests(path string) ([]*protoCommand.CommandRequest, error) {
	var requests []*protoCommand.CommandRequest
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &requests)
	if err != nil {
		return nil, err
	}
	return requests, nil
}
