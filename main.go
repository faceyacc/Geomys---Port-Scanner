package main

import (
	"fmt"

	"github.com/faceyacc/geomys/port"
)

func main() {

	fmt.Println("geomys - port scanner")
	scanTCP, scanUDP := port.InitalScan("localhost")

	fmt.Println(scanTCP)
	fmt.Println(scanUDP)

}
