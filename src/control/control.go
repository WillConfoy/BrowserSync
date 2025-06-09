package control

import (
	"fmt"
	"log"

	"strings"

	rb "github.com/go-vgo/robotgo"
	gohook "github.com/robotn/gohook"

	node "cs498.com/browsersync/node"
	s "cs498.com/browsersync/structs"
)

var (
	window      string // window is the string that needs to be in the name of the window for any inputs to be accepted
	sb          strings.Builder
	pressed     map[uint16]bool   = make(map[uint16]bool)
	rawcodedict map[uint16]string = map[uint16]string{
		162: "l ctrl", 163: "r ctrl", 164: "l alt", 165: "r alt", 91: "win", 160: "l shift", 161: "r shift", 9: "tab", 13: "enter",
		8: "backspace", 27: "esc", 32: "space"}
)

func Start(startingstate *s.StateInfo, machine *s.MachineInfo) {
	mynode := node.Start(startingstate, machine)
	go mynode.Run(startingstate, machine) // this is just starting up the node
	log.Printf("Current PID: %d\n", rb.GetPid())

	// We ignore any inputs not in a window that has "main.go" somewhere in the title. Practically, I'd set this to "google chrome"
	handleInputs(startingstate, &mynode) // we now go into an infinite loop so that we can handle events as the come up

}

// This function ensures that we are only sending and receiving inputs that were made in the correct window
func CheckRightWindow(window string) bool {
	active := strings.ToLower(rb.GetTitle())
	log.Println("Current active window: ", active)
	return strings.Contains(active, window)
}

// This function loops infinitely waiting for keystrokes and mouse events from the user
func handleInputs(startingstate *s.StateInfo, mynode *node.Node) {
	// TODO: Find a way to ignore inputs from robotgo while paying attention to ones that are from the user
	// If not possible, stick with leader design like we have here
	// gohook.AddEvents("q", "ctrl", "shift")
	if mynode.Leader {
		fmt.Println("You are the leader")
	} else {
		fmt.Println("You are not the leader")
	}

	if startingstate.Allowtransfer {
		fmt.Println("Anyone can become the leader by pressing f9")
	}

	eventHook := gohook.Start()
	var e gohook.Event

	for e = range eventHook {

		if !mynode.Leader && e.Rawcode != 120 {
			continue
		} else if !mynode.Leader && startingstate.Allowtransfer && e.Kind == gohook.KeyUp {
			log.Println("BECOMING LEADER!!!")
			log.Println(e)
			mynode.Leader = true
			mynode.BroadcastNewLeader()
			continue
		}

		if e.Kind == gohook.MouseDown {
			if !CheckRightWindow(window) {
				continue
			}
			handleClicks(e, mynode)
		} else if e.Kind == gohook.KeyDown {
			if pressed[e.Rawcode] || !CheckRightWindow(window) || !mynode.Leader {
				continue
			}
			pressed[e.Rawcode] = true

			// if !handleCommand(e, mynode) {
			handleKeydown(e, mynode)
			// }
		} else if e.Kind == gohook.KeyHold {
			log.Println(e)
			if pressed[e.Rawcode] || !CheckRightWindow(window) {
				continue
			}

			pressed[e.Rawcode] = true
			if !handleCommand(e, mynode) {
				// handleKeydown(e, mynode)
				handleKeyholds(e, mynode)
			}

		} else if e.Kind == gohook.KeyUp {
			if !CheckRightWindow(window) {
				continue
			}
			handleKeyups(e, mynode)
			pressed[e.Rawcode] = false
		} else if e.Kind == gohook.MouseWheel {
			if !CheckRightWindow(window) {
				continue
			}
			handleScrolls(e, mynode)
		}
	}
}

//
// All the functions that start with 'handle' take care of telling the node to send the appropriate signal across the network
//

func handleClicks(e gohook.Event, mynode *node.Node) {
	mynode.SendClick(int(e.Button), int(e.X), int(e.Y))
	log.Printf("Clicked Button %d -- X: %d, Y: %d\n", e.Button, e.X, e.Y)
	if e.Button == 3 {
		log.Println(sb.String())
	}
}

func handleKeydown(e gohook.Event, mynode *node.Node) {

	mynode.SendKeyDown(string(e.Keychar))

	sb.WriteString(string(e.Keychar))
	log.Printf("Pressed key down: %s\n", string(e.Keychar))
}

// Get all held keys and send them all at once
func handleCommand(e gohook.Event, mynode *node.Node) bool {
	log.Println("GONNA TRY SENDING A COMMAND!!!")
	delim_string := ""
	// construct delimited string of all pressed keys
	for key, b := range pressed {
		if b {
			r, is_in := rawcodedict[key]

			if is_in {
				delim_string += (r + "|")
			} else {
				delim_string += string(gohook.RawcodetoKeychar(e.Rawcode)) + "|"
			}
		}
	}

	if len(delim_string) == 0 {
		return false
	}

	log.Println("DELIMITED STRING OF HELD KEYS IS: " + delim_string)

	r, is_in := rawcodedict[e.Rawcode]
	if is_in {
		delim_string += r
		mynode.SendKeyDown(r)
	} else {
		delim_string += string(gohook.RawcodetoKeychar(e.Rawcode))
	}

	log.Println("FINAL DELIMITED STRING IS: " + delim_string)
	// log.Printf("Key command: %s \n", delim_string)
	sb.WriteString(delim_string)
	mynode.SendCommand(delim_string)

	return true
}

func handleKeyholds(e gohook.Event, mynode *node.Node) {
	r, is_in := rawcodedict[e.Rawcode]
	if is_in {
		sb.WriteString(r)
		log.Printf("Key held: %s \n", r)

		mynode.SendKeyDown(r)
	} else {
		newstr := string(gohook.RawcodetoKeychar(e.Rawcode))
		sb.WriteString(newstr)
		log.Printf("Key held: %s \n", newstr)

		mynode.SendKeyDown(newstr)
	}
}

func handleKeyups(e gohook.Event, mynode *node.Node) {
	r, is_in := rawcodedict[e.Rawcode]
	if is_in {
		// sb.WriteString(r)
		log.Printf("Released key: %s\n", r)

		mynode.SendKeyUp(r)
	} else {
		gohook.RawcodetoKeychar(e.Rawcode)
		newstr := string(gohook.RawcodetoKeychar(e.Rawcode))
		log.Printf("Released key: %s\n", newstr)

		mynode.SendKeyUp(newstr)
	}
}

func handleScrolls(e gohook.Event, mynode *node.Node) {
	if e.Rotation == 1 {
		mynode.SendScroll("down")
	} else {
		mynode.SendScroll("up")
	}
}
