package main

import (
	"flag"

	control "cs498.com/browsersync/control"
	gather "cs498.com/browsersync/gather"
)

var (
	manual        = flag.Bool("manual", false, "Whether to enter args manually or use website integration")
	leader        = flag.Bool("l", false, "Whether or not this node is the leader")
	allowtransfer = flag.Bool("allow", true, "Whether people can take leader")
	addrstring    = flag.String("addrs", "localhost:50051", "A space separated list of addresses with ports. Ex: 127.0.0.1:50051")
	myport        = flag.String("port", "50051", "The port for this node to run on")
	myip          = flag.String("myip", "-1", "The ip to run on")
	thewindow     = flag.String("w", "google chrome", "Part of the name of the window ex: google chrome")
)

func main() {
	flag.Parse()
	if *manual {
		control.Start(*leader, *allowtransfer, *addrstring, *myport, *myip, *thewindow)
	} else {
		*leader, *allowtransfer, *addrstring, *myport, *myip, *thewindow = gather.Gather()
		control.Start(*leader, *allowtransfer, *addrstring, *myport, *myip, *thewindow)
	}
}
