package port

import (
	"net"
	"strconv"
	"time"
)

type Scan struct {
	Port    string
	State   string
	Service string
}

func ScanPort(protocol, hostname string, port int) Scan {
	scan := Scan{Port: strconv.Itoa(port) + "/" + protocol}
	address := hostname + ":" + strconv.Itoa(port)
	connection, err := net.DialTimeout(protocol, address, 60*time.Second) // Connect to address

	if err != nil {
		scan.State = "Closed"
		return scan
	}

	defer connection.Close()
	scan.State = "Open"
	return scan
}

func InitalScan(hostname string) ([]Scan, []Scan) {

	var resultTCP []Scan
	var resultUDP []Scan

	for i := 0; i < 1024; i++ {
		resultTCP = append(resultTCP, ScanPort("tcp", hostname, i))
		resultUDP = append(resultUDP, ScanPort("udp", hostname, i))
	}

	return resultTCP, resultUDP
}
