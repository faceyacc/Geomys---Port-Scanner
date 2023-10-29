package port

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/tatsushid/go-fastping"
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

func ICMPScan(host string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		fmt.Println("finish")
	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func WideScan(hostname string) ([]Scan, []Scan) {
	var resultTCP []Scan
	var resultUDP []Scan

	for i := 1024; i < 49152; i++ {
		resultTCP = append(resultTCP, ScanPort("tcp", hostname, i))
		resultUDP = append(resultUDP, ScanPort("udp", hostname, i))
	}
	return resultTCP, resultUDP
}
