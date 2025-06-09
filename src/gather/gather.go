package gather

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
	"time"

	"google.golang.org/grpc"

	gs "cs498.com/browsersync/gatherproto"
	s "cs498.com/browsersync/structs"
)

var (
	stopchan      chan int
	machine       s.MachineInfo
	startingstate s.StateInfo
	port          = "20601"
)

type GatherServer struct {
	gs.UnimplementedGatherServiceServer
}

func Gather() (s.StateInfo, s.MachineInfo, error) {
	fmt.Println("Hi!!!")
	startListening()
	if !checkState() {
		return startingstate, machine, errors.New("unable to get startingstate/machine values")
	}
	return startingstate, machine, nil
}

func startListening() {
	app := "tailscale"
	arg0 := "ip"
	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to get tailscale ip: %v", err)
	}
	ip := strings.Split(string(stdout), "\n")[0]

	lis, err := net.Listen("tcp", ip+":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer() // n is for serving purpose
	myserver := &GatherServer{}

	gs.RegisterGatherServiceServer(grpcServer, myserver)

	go func() {
		<-stopchan
		time.Sleep(1 * time.Second)
		grpcServer.GracefulStop()
	}()

	// start listening
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func checkState() bool {
	return machine.Initialized && startingstate.Initialized
}

func (gather *GatherServer) EndServer(ctx context.Context, in *gs.EndServerRequest) (*gs.EndServerResponse, error) {
	stopchan <- 1
	return &gs.EndServerResponse{Success: true}, nil
}

func (gather *GatherServer) SendMachineInfo(ctx context.Context, in *gs.MachineInfoRequest) (*gs.MachineInfoResponse, error) {
	machine = s.MachineInfo{Port: in.Port, Ip: in.Ip, Window: in.Window, Initialized: true}
	return &gs.MachineInfoResponse{Success: true}, nil
}

func (gather *GatherServer) SendStateInfo(ctx context.Context, in *gs.StateInfoRequest) (*gs.StateInfoResponse, error) {
	addrs := strings.Split(in.Addrstring, "|")
	startingstate = s.StateInfo{Leader: in.Leader, Allowtransfer: in.Allowtransfer, Addrs: addrs, Initialized: true}
	return &gs.StateInfoResponse{Success: true}, nil
}
