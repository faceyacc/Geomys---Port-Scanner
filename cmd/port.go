package cmd

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/Ullaakut/nmap/v3"
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

func ICMPScan(hostname string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", hostname)
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

func NMAPScan(hostname, port string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		ctx,
		nmap.WithTargets(hostname),
		nmap.WithPorts(port),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, warnings, err := scanner.Run()
	if len(*warnings) > 0 {
		log.Printf("run finished with warnings: %s\n", warnings)
	}
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Scan completed: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}
