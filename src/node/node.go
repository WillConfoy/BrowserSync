package node

import (
	"context"
	// "errors"
	// "fmt"
	"log"
	"slices"
	"strings"

	// "math/rand"
	// "net/http"
	// "path/filepath"
	// "sync"
	"net"
	"time"

	"google.golang.org/grpc"

	// "google.golang.org/grpc/reflection"
	"google.golang.org/grpc/credentials/insecure"
	// gohook "github.com/robotn/gohook"
	rs "cs498.com/browsersync/nodeproto"
	s "cs498.com/browsersync/structs"
	rgo "github.com/go-vgo/robotgo"
)

var (
	// sb strings.Builder
	// counter int = 0
	// pressed map[uint16]bool = make(map[uint16]bool)
	// rawcodedict map[uint16]string = map[uint16]string{
	//     162:"l ctrl", 163:"r ctrl", 164:"l alt", 165:"r alt", 91:"win", 160:"l shift", 161:"r shift", 9:"tab", 13:"enter",
	//     8:"backspace", 27:"esc", 32:"space"}
	buttonmap map[int]string = map[int]string{1: "left", 2: "right", 3: "center"}
)

// In order, we have Leader- whether or not the node is currently the leader, currentleader- the IP of the current leader, peers- a map of IPs to clients,
// myaddr- the IP of the node, and the final thing is just needed for gRPC to work
type Node struct {
	Leader        bool
	currentleader string
	Peers         map[string]rs.SyncServiceClient
	myaddr        string
	rs.UnimplementedSyncServiceServer
	maxX   int
	maxY   int
	Window string
}

// This function is the most basic initialization of the node
func Start(startingstate *s.StateInfo, machine *s.MachineInfo) Node {
	log.Println("HI IT'S ME THE NODE!!!")
	x, y := rgo.GetScreenSize()
	// mynode := Node{Leader: leader, maxX: rgo.GetScreenRect().W, maxY: rgo.GetScreenRect().H, Window: window}

	mynode := Node{Leader: startingstate.Leader, maxX: x, maxY: y, Window: machine.Window}
	return mynode
}

// This function does the actual running- it initializes the fields of the node before listening and starting up the heartbeat
func (node *Node) Run(startingstate *s.StateInfo, machine *s.MachineInfo) {
	// node.myaddr = "localhost:" + port
	if machine.Ip == "-1" {
		node.myaddr = GetLocalIP() + ":" + machine.Port
	} else {
		node.myaddr = machine.Ip + ":" + machine.Port
	}

	log.Println(node.myaddr)
	startingstate.Addrs = slices.DeleteFunc(startingstate.Addrs, func(addr string) bool { return node.myaddr == addr })
	node.Peers = make(map[string]rs.SyncServiceClient)

	go node.StartListening()

	time.Sleep(5 * time.Second)
	node.createPeers(startingstate.Addrs)
	log.Println("HI I AM RUNNING")

	log.Println(node.Peers)
	time.Sleep(time.Second)

	if node.Leader {
		node.currentleader = node.myaddr
		node.BroadcastNewLeader()
	}

	node.DoHeartbeat()
}

// Got this directly from stackoverflow
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// Make sure we're in the right window
func CheckRightWindow(window string) bool {
	active := strings.ToLower(rgo.GetTitle())
	return strings.Contains(active, window)
}

// This function gets called in another goroutine and starts listening
func (node *Node) StartListening() {
	lis, err := net.Listen("tcp", node.myaddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer() // n is for serving purpose

	rs.RegisterSyncServiceServer(grpcServer, node)

	// start listening
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// This function creates a client for every peer listed in the array of peer IPs
func (node *Node) createPeers(addrs []string) {
	for _, x := range addrs {
		node.createClient(x)
	}
}

// This creates a client and puts it into the peer map
func (node *Node) createClient(addr string) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error connecting to %s, error %v", addr, err)
	}

	// defer conn.Close()

	node.Peers[addr] = rs.NewSyncServiceClient(conn)

	response, err := node.Peers[addr].HeartbeatInternal(context.Background(), &rs.HeartbeatRequest{Beat: node.myaddr})
	if err != nil {
		log.Printf("Error connecting to %s, error %v", addr, err)
	}

	log.Printf("Greeting from other node: %s", response.GetRet())
}

// This is just a test function to make sure the node exists and it kind of outdated
func (node *Node) PrintStuff() {
	log.Println("I am printing stuff")
}

// Once per second, sends heartbeat out
func (node *Node) DoHeartbeat() {
	for {
		time.Sleep(time.Second)
		node.BroadcastHeartbeat()
		// log.Println(node.Peers)
	}
}

// Sends a heartbeat signal to every peer
func (node *Node) BroadcastHeartbeat() {
	for ip, client := range node.Peers {
		r, err := client.HeartbeatInternal(context.Background(), &rs.HeartbeatRequest{Beat: node.myaddr})

		if err != nil {
			log.Printf("HEARTBEAT TO %s FAILED\n", ip)
		} else if r.GetRet() == node.myaddr {
			// log.Printf("GOT HEARTBEAT FROM: %s\n", ip)
		} else {
			log.Fatalf("SOMETHING WENT WRONG AND %s RESPONDED WITH %s INSTEAD OF %s\n", ip, r.GetRet(), node.myaddr)
		}
	}
}

// This will call UpdateLeader on every peer
func (node *Node) BroadcastNewLeader() {
	for ip, client := range node.Peers {
		_, err := client.UpdateLeader(context.Background(), &rs.LeaderRequest{Ip: node.myaddr})

		if err != nil {
			log.Printf("FAILED SENDING NEW LEADER %s TO %s\n", node.myaddr, ip)
		} else {
			log.Printf("Updated New Leader: %s for %s\n", node.myaddr, ip)
		}
	}
}

// Sends a signal to click a mouse button to every peer
func (node *Node) SendClick(button int, x int, y int) {
	buttonString := buttonmap[button]
	newX := float64(x) / float64(node.maxX)
	newY := float64(y) / float64(node.maxY)

	for ip, client := range node.Peers {
		response, err := client.SendClickInternal(context.Background(), &rs.ClickRequest{
			Button:   buttonString,
			XPercent: newX,
			YPercent: newY,
		})

		if err != nil {
			log.Printf("FAILED SENDING %s CLICK TO %s\n", buttonString, ip)
		} else {
			log.Printf("Sent Click: %s to %s: %t\n", buttonString, ip, response.GetSuccess())
		}
	}
}

// Sends a signal to hold a key down to every peer
func (node *Node) SendKeyDown(key string) {
	for ip, client := range node.Peers {
		response, err := client.SendKeyDownInternal(context.Background(), &rs.KeyDownRequest{
			Key: key,
		})

		if err != nil {
			log.Printf("FAILED SENDING %s KEY TO %s\n", key, ip)
		} else {
			log.Printf("Sent Key: %s to %s: %t\n", key, ip, response.GetSuccess())
		}
	}
}

func (node *Node) SendCommand(keys string) {
	for ip, client := range node.Peers {
		response, err := client.SendKeyUpInternal(context.Background(), &rs.KeyUpRequest{
			Key: keys,
		})

		if err != nil {
			log.Printf("FAILED SENDING %s KEY TO %s\n", keys, ip)
		} else {
			log.Printf("Sent COMMAND Keys: %s to %s: %t\n", keys, ip, response.GetSuccess())
		}
	}
}

// Sends a signal to release a key to every peer
func (node *Node) SendKeyUp(key string) {
	for ip, client := range node.Peers {
		response, err := client.SendKeyUpInternal(context.Background(), &rs.KeyUpRequest{
			Key: key,
		})

		if err != nil {
			log.Printf("FAILED SENDING %s KEY TO %s\n", key, ip)
		} else {
			log.Printf("Sent Key: %s to %s: %t\n", key, ip, response.GetSuccess())
		}
	}
}

// Sends a signal to every peer to scroll a certain direction
func (node *Node) SendScroll(direction string) {
	for ip, client := range node.Peers {
		_, err := client.SendScrollInternal(context.Background(), &rs.ScrollRequest{
			Direction: direction,
		})

		if err != nil {
			log.Printf("FAILED SENDING '%s' SCROLL TO %s\n", direction, ip)
		} else {
			// log.Printf("Sent Scroll: '%s' to %s: %t\n", direction, ip, response.GetSuccess())
		}
	}
}

//
// //
// // // THESE FOLLOWING 5 FUNCTIONS ACTUALLY DO THE INPUT USING ROBOTGO, LIKE CLICKING THE MOUSE AT A CERTAIN X AND Y COORDINATE
// //
//

func (node *Node) HeartbeatInternal(ctx context.Context, in *rs.HeartbeatRequest) (*rs.HeartbeatResponse, error) {
	return &rs.HeartbeatResponse{Ret: in.GetBeat()}, nil
}

func (node *Node) SendClickInternal(ctx context.Context, in *rs.ClickRequest) (*rs.ClickResponse, error) {
	x, y := rgo.GetScreenSize()
	log.Printf("Max X: %d, Max Y: %d, Percent X: %f, Percent Y: %f", node.maxX, node.maxY, in.GetXPercent(), in.GetYPercent())
	log.Printf("New Max X: %d, New Max Y: %d", x, y)
	if CheckRightWindow(node.Window) {
		newX := int(in.GetXPercent() * float64(node.maxX))
		newY := int(in.GetYPercent() * float64(node.maxY))
		rgo.MoveClick(newX, newY, in.GetButton())
		return &rs.ClickResponse{Success: true}, nil
	} else {
		log.Printf("Not in right window- current window: %s, desired string: %s\n", strings.ToLower(rgo.GetTitle()), node.Window)
		return &rs.ClickResponse{Success: false}, nil
	}
}

func (node *Node) SendKeyDownInternal(ctx context.Context, in *rs.KeyDownRequest) (*rs.KeyDownResponse, error) {
	if CheckRightWindow(node.Window) {
		rgo.KeyDown(in.GetKey())
		return &rs.KeyDownResponse{Success: true}, nil
	} else {
		log.Printf("Not in right window- current window: %s, desired string: %s\n", strings.ToLower(rgo.GetTitle()), node.Window)
		return &rs.KeyDownResponse{Success: false}, nil
	}
}

func (node *Node) SendCommandInternal(ctx context.Context, in *rs.CommandRequest) (*rs.CommandResponse, error) {
	if CheckRightWindow(node.Window) {
		keys := strings.Split(in.GetCommand(), "|")
		rgo.KeyTap(keys[0], keys[1:])
		log.Println("GOT COMMAND!!! TRYING TO PRESS THE KEYS " + strings.Join(keys, ","))
		return &rs.CommandResponse{Success: true}, nil
	} else {
		log.Printf("Not in right window- current window: %s, desired string: %s\n", strings.ToLower(rgo.GetTitle()), node.Window)
		return &rs.CommandResponse{Success: false}, nil
	}
}

func (node *Node) SendKeyUpInternal(ctx context.Context, in *rs.KeyUpRequest) (*rs.KeyUpResponse, error) {
	if CheckRightWindow(node.Window) {
		rgo.KeyUp(in.GetKey())
		return &rs.KeyUpResponse{Success: true}, nil
	} else {
		log.Printf("Not in right window- current window: %s, desired string: %s\n", strings.ToLower(rgo.GetTitle()), node.Window)
		return &rs.KeyUpResponse{Success: false}, nil
	}
}

func (node *Node) SendScrollInternal(ctx context.Context, in *rs.ScrollRequest) (*rs.ScrollResponse, error) {
	if CheckRightWindow(node.Window) {
		rgo.ScrollDir(1, in.GetDirection())
		return &rs.ScrollResponse{Success: true}, nil
	} else {
		log.Printf("Not in right window- current window: %s, desired string: %s\n", strings.ToLower(rgo.GetTitle()), node.Window)
		return &rs.ScrollResponse{Success: false}, nil
	}
}

// This works by having the requesting node call this on every other machine
func (node *Node) UpdateLeader(ctx context.Context, in *rs.LeaderRequest) (*rs.LeaderResponse, error) {
	log.Printf("Leader becoming false...\n")
	node.currentleader = in.GetIp()
	node.Leader = false
	log.Printf("Current state - Leader ip: %s, Leader bool: %t\n", node.currentleader, node.Leader)
	// if node.currentleader == node.myaddr {node.Leader = false}
	return &rs.LeaderResponse{}, nil
}
