package server

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/mittachaitu/Learning/grpc-k8s/protobuffers/go/command"
	"google.golang.org/grpc"
)

type Server struct {
	rootDir     string
	jailCommand string
	ipAddress   string
	port        string
	command.UnimplementedRunRPCCommandsServer
}

func NewServer(ipAddress, port string) *Server {
	return &Server{
		rootDir:     os.Getenv("RootDirectory"),
		jailCommand: "chroot",
		ipAddress:   ipAddress,
		port:        port,
	}
}

func (s *Server) Start() error {
	fmt.Println("Creating gRPC server")
	lis, err := net.Listen("tcp4", s.ipAddress+":"+s.port)
	if err != nil {
		return fmt.Errorf("Failed to listen on address: %s error: %v", s.ipAddress+":"+s.port, err)
	}
	grpcServer := grpc.NewServer()
	command.RegisterRunRPCCommandsServer(grpcServer, s)
	if err = grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("Failed to start server error: %v", err)
	}
	return nil
}

func (s *Server) RunShellCommands(ctx context.Context, in *command.ShellCommandsOutputRequest) (*command.ShellCommandsOutputResponse, error) {
	responses, err := s.executeShellCommands(in.Requests)
	if err != nil {
		return nil, err
	}
	return &command.ShellCommandsOutputResponse{Responses: responses}, nil
}

func (s *Server) executeShellCommands(commandRequests []*command.CommandRequest) ([]*command.CommandResponse, error) {
	var result []*command.CommandResponse
	result = make([]*command.CommandResponse, len(commandRequests))
	for i := 0; i < len(commandRequests); i++ {
		output, err := s.executeCommand(commandRequests[i].Binary, commandRequests[i].Args)
		if err != nil {
			return nil, err
		}
		result[i] = &command.CommandResponse{
			Output: output,
		}
	}
	return result, nil
}

func (s *Server) executeCommand(bin string, args []string) ([]byte, error) {
	var jailCommand, commandBuilder strings.Builder
	commandBuilder.WriteString(bin)
	for _, arg := range args {
		commandBuilder.WriteString(" " + arg)
	}

	if s.rootDir != "" {
		jailCommand.WriteString(s.jailCommand)
		jailCommand.WriteString(" " + s.rootDir)
		jailCommand.WriteString(" /bin/sh")
		jailCommand.WriteString(" -c")
		jailCommand.WriteString(" \"")
		jailCommand.WriteString(commandBuilder.String())
		jailCommand.WriteString("\"")
		commandBuilder = jailCommand
	}
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("/bin/sh", "-c", commandBuilder.String())
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to run command: %q error: %v stderror: %s", commandBuilder.String(), err, string(stderr.Bytes()))
	}
	return stdout.Bytes(), nil
}
