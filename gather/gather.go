package gather

import (
	"flag"
	"fmt"
)

var ()

func Gather() (bool, bool, string, string, string, string) {
	flag.Parse()
	fmt.Println("Hi!!!")
	return true, true, "", "", "", ""
}
