package gather

import (
	"flag"
	"fmt"
	"strings"

	s "cs498.com/browsersync/structs"
)

var ()

func Gather() (s.StateInfo, s.MachineInfo) {
	flag.Parse()
	fmt.Println("Hi!!!")
	state := s.StateInfo{Leader: true, Allowtransfer: true, Addrs: strings.Split("hi|me", "|")}
	machine := s.MachineInfo{Port: "100", Ip: "127.0.0.1", Window: "firefox"}
	return state, machine
}
