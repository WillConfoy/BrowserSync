package structs

type StateInfo struct {
	Leader        bool
	Allowtransfer bool
	Addrs         []string
	Initialized   bool
}

type MachineInfo struct {
	Port        string
	Ip          string
	Window      string
	Initialized bool
	Displaynum  int
}
