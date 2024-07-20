package main

import (
	// "context"
	// "errors"
	"fmt"
	"log"
	// "time"

	// "os"
	"flag"
	"strings"

	// "math/rand"
	// "net/http"
	// "path/filepath"
	// "sync"
	// "time"

	rb "github.com/go-vgo/robotgo"
	gohook "github.com/robotn/gohook"

	node "cs498.com/browsersync/node"
)

// // This example opens https://github.com/, searches for "git",
// // and then gets the header element which gives the description for Git.
// func main() {
// 	path, _ := launcher.LookPath()

// 	launch := launcher.New().Leakless(false)
//     u := launch.
// 		Bin(path).
// 		MustLaunch()

// 	defer launch.Cleanup()

// 	page := rod.New().ControlURL(u).MustPage("https://www.youtube.com/")
// 	fmt.Printf("HI\n")
// 	// browser := rod.New().ControlURL(u)
// 	// page := browser.MustPage("https://www.youtube.com/")
// 	// browser.Page(proto.TargetCreateTarget{URL: "http://www.wikipedia.org"})
// 	// page := rod.New().ControlURL(u).MustConnect().MustPage("https://www.wikipedia.org/")
//     // page.MustWindowFullscreen()

// 	// go page.EachEvent(func(e *proto.RuntimeConsoleAPICalled) {
// 	// 	fmt.Println(page.MustObjectsToJSON(e.Args))
// 	// })()

// 	// page.WaitEvent(proto.)

// 	// rod.Mouse.page.e()
// 	time.Sleep(4 * time.Second)
//     // page.MustWaitStable().MustScreenshot("a.png")
//     page.MustScreenshot("a.png")
//     time.Sleep(time.Hour)
// }

var (
    window string   // window is the string that needs to be in the name of the window for any inputs to be accepted
    sb strings.Builder
    pressed map[uint16]bool = make(map[uint16]bool)
    rawcodedict map[uint16]string = map[uint16]string{
        162:"l ctrl", 163:"r ctrl", 164:"l alt", 165:"r alt", 91:"win", 160:"l shift", 161:"r shift", 9:"tab", 13:"enter",
        8:"backspace", 27:"esc", 32:"space"}
    leader = flag.Bool("l", false, "Whether or not this node is the leader")
    allowtransfer bool = true
    addrstring = flag.String("addrs", "localhost:50051", "A space separated list of addresses with ports. Ex: 127.0.0.1:50051")
	myport = flag.String("port", "50051", "The port for this node to run on")
    myip = flag.String("myip", "-1", "The ip to run on")
    thewindow = flag.String("w", "google chrome", "Part of the name of the window ex: google chrome")
)

func main() {
    flag.Parse()
    addrs := strings.Split(*addrstring, " ")
    window = *thewindow  // So this is saying we ignore any inputs not in a window that has "main.go" somewhere in the title. Practically, I'd set this to "google chrome"
    mynode := node.Start(*leader, window)
    go mynode.Run(addrs, *myip, *myport)
    // mynode.PrintStuff()aa
    // log.Println(mynode.Leader)aaaaaaaaaa
    // os.Exit(0)
    log.Printf("Current PID: %d\n", rb.GetPid())

    handleInputs(window, &mynode)


    // tempnode := node.Node{Leader: false}
    // handleInputs(window, tempnode)

    // for {
    //     time.Sleep(2 * time.Second)
    //     // rb.KeyDown("w")
    //     time.Sleep(50 * time.Millisecond)
    //     // rb.KeyUp("w")
    // }
}

func CheckRightWindow(window string) bool {
    active := strings.ToLower(rb.GetTitle())
    log.Println("Current active window: ", active)
    return strings.Contains(active, window)
}


func handleInputs(window string, mynode *node.Node) {
    // TODO: Find a way to ignore inputs from robotgo while paying attention to ones that are from the user
    // If not possible, stick with leader design like we have here
    // gohook.AddEvents("q", "ctrl", "shift")
    if mynode.Leader {
        fmt.Println("You are the leader")
    } else {
        fmt.Println("You are not the leader")
    }

    if allowtransfer {fmt.Println("Anyone can become the leader by pressing f9")}

    eventHook := gohook.Start()
    var e gohook.Event

    // for ok := true; ok; ok = mynode.Leader {
    // for {
        for e = range eventHook {

            // for !mynode.Leader {
            //     time.Sleep(10 * time.Second)
            // }

            if !mynode.Leader && e.Rawcode != 120 {
                continue
            } else if !mynode.Leader && allowtransfer && e.Kind != gohook.KeyUp {
                log.Println("BECOMING LEADER!!!")
                log.Println(e)
                mynode.Leader = true
                mynode.BroadcastNewLeader()
                continue
            }

            // if !mynode.Leader {break}

            if e.Kind == gohook.MouseDown {
                // log.Println(e)
                if !CheckRightWindow(window) {continue}
                handleClicks(e, mynode)
            } else if e.Kind == gohook.KeyDown {
                // log.Println(e)
                // log.Printf("Current pressed value: %t\n", pressed[e.Rawcode])
                if pressed[e.Rawcode] || !CheckRightWindow(window) || !*leader {continue}
                // if !CheckRightWindow(window) {continue}
                // if !pressed[string(e.Keychar)] {continue}
                pressed[e.Rawcode] = true
                handleKeydown(e, mynode)
            } else if e.Kind == gohook.KeyHold {
                log.Println(e)
                if pressed[e.Rawcode] || !CheckRightWindow(window) {continue}

                pressed[e.Rawcode] = true
                handleKeyholds(e, mynode)

            } else if e.Kind == gohook.KeyUp {
                // log.Println(e)
                if !CheckRightWindow(window) {continue}

                handleKeyups(e, mynode)
                pressed[e.Rawcode] = false

                // log.Printf("Current pressed value: %t\n", pressed[e.Rawcode])
                // log.Println(pressed)
            } else if e.Kind == gohook.MouseWheel {
                // log.Println(e)
                if !CheckRightWindow(window) {continue}
                handleScrolls(e, mynode)
            }
        }
    //     log.Printf("My value should be false: %t\n", mynode.Leader)
    //     time.Sleep(1 * time.Second)
    // }
}

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
        // log.Printf("Released key: %s\n", string(e.Keychar))
        newstr := string(gohook.RawcodetoKeychar(e.Rawcode))
        // sb.WriteString(newstr)
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


// test: Hello WorldだんしゃりHi, Seattle space needle, Golden gate bridge, One world trade center.Hi galaxy, hi stars, hi MT.Rainier, hi sea. こんにちは世界.
// a: Hello WorldだんしゃりHi, Seattle space needle, Golden gate bridge, One world trade center.Hi galaxy, hi stars, hi MT.Rainier, hi sea. こんにちは世界.a
// test: Hello WorldだんしゃりHi, Seattle space needle, Golden gate bridge, One world trade center.Hi galaxy, hi stars, hi MT.Rainier, hi sea. こんにちは世界.w
// eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeewqwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwweeewwwwwwwwwwwww
// wwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww

// wwwwwwwwwwwww This is a cool string that should be soon doneHopefully this is correctwwwwww