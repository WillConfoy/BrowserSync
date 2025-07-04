package main

import (
	"embed"
	"flag"
	"log"
	"strings"
	"time"

	control "cs498.com/browsersync/control"
	gather "cs498.com/browsersync/gather"
	s "cs498.com/browsersync/structs"
	"github.com/gopxl/beep/v2"
	"github.com/gopxl/beep/v2/mp3"
	"github.com/gopxl/beep/v2/speaker"
)

var (
	manual        = flag.Bool("manual", false, "Whether to enter args manually or use website integration")
	leader        = flag.Bool("l", false, "Original leader")
	allowtransfer = flag.Bool("allow", true, "Whether people can take leader")
	addrstring    = flag.String("addrs", "localhost:50051", "A space separated list of addresses with ports. Ex: 127.0.0.1:50051")
	myport        = flag.String("port", "50051", "The port for this node to run on")
	myip          = flag.String("myip", "-1", "The ip to run on")
	thewindow     = flag.String("w", "google chrome", "Part of the name of the window ex: google chrome")
)

//go:embed alerts
var alertsFS embed.FS

func main() {
	flag.Parse()
	var startingstate s.StateInfo
	var machine s.MachineInfo
	var err error

	if *manual {
		startingstate = s.StateInfo{Leader: *leader, Allowtransfer: *allowtransfer, Addrs: strings.Split(*addrstring, " "), Initialized: true}
		machine = s.MachineInfo{Port: *myport, Ip: *myip, Window: *thewindow, Initialized: true}
	} else {
		startingstate, machine, err = gather.Gather()
	}

	// make sure to change this before actually using this code bc right now it does error always
	if err != nil {
		log.Println("Error getting state! Please speak up and tell the others in the session!")

		f, err := alertsFS.Open("alerts/error_alert.mp3")
		if err != nil {
			log.Fatal(err)
		}
		streamer, format, err := mp3.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		// Play error_alert.mp3
		speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		done := make(chan bool)
		speaker.Play(beep.Seq(streamer, beep.Callback(func() {
			done <- true
		})))
		<-done
	} else {
		control.Start(&startingstate, &machine)
	}
}
