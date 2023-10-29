package main

import (
	"fmt"

	"github.com/faceyacc/geomys/port"
)

func main() {

	fmt.Println("geomys - port scanner")
	// port.ICMPScan("localhost")
	// fmt.Println(test)
	scanTCP, scanUDP := port.WideScan("localhost")
	fmt.Println(scanTCP)
	fmt.Println(scanUDP)

}
