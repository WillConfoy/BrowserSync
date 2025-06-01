package structs

type StateInfo struct {
	Leader        bool
	Allowtransfer bool
	Addrs         []string
}

type MachineInfo struct {
	Port   string
	Ip     string
	Window string
}
