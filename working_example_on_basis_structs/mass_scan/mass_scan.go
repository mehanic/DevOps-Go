package main

import (
	"encoding/xml"
	"fmt"
	"os"
//	"path/filepath"
//	"sort"
	"strings"
)

// HostItem holds information about a host and its open ports.
type HostItem struct {
	IP    string
	Ports []PortItem
}

// PortItem holds information about an open port on a host.
type PortItem struct {
	Port    int
	Proto   string
	Service string
	Banner  string
}

// XML structures to match Masscan output.
type NmapRun struct {
	Hosts []Host `xml:"host"`
}

type Host struct {
	Address string `xml:"address,attr"`
	Ports   []Port `xml:"ports>port"`
}

type Port struct {
	PortID  int    `xml:"portid,attr"`
	Proto   string `xml:"protocol,attr"`
	Service Service `xml:"service"`
}

type Service struct {
	Name    string `xml:"name,attr"`
	Product string `xml:"product,attr"`
	Banner  string `xml:"banner,attr"`
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	fileName := os.Args[1]

	nmapRun, err := openMasscanFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Process each host and print the details
	for _, host := range nmapRun.Hosts {
		h := createHost(host)
		fmt.Println(h)
	}
}

func usage() {
	fmt.Println("Usage: masscan_parse masscan_xml_file")
	os.Exit(1)
}

func openMasscanFile(fileName string) (*NmapRun, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s does not exist", fileName)
	}

	if !isFile(fileName) {
		return nil, fmt.Errorf("%s is not a file", fileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	var nmapRun NmapRun
	if err := xml.NewDecoder(file).Decode(&nmapRun); err != nil {
		return nil, fmt.Errorf("XML Parse Error: %s", err)
	}

	// Validate if it is a Masscan XML file by checking attributes
	if len(nmapRun.Hosts) == 0 {
		return nil, fmt.Errorf("%s is not a Masscan XML file", fileName)
	}

	return &nmapRun, nil
}

func isFile(fileName string) bool {
	info, err := os.Stat(fileName)
	return err == nil && !info.IsDir()
}

func createHost(host Host) HostItem {
	h := HostItem{
		IP: host.Address,
	}

	for _, port := range host.Ports {
		serviceName := getServiceName(port.Service)
		h.Ports = append(h.Ports, PortItem{
			Port:    port.PortID,
			Proto:   port.Proto,
			Service: serviceName,
			Banner:  port.Service.Banner,
		})
	}

	return h
}

func getServiceName(service Service) string {
	if service.Product != "" {
		return service.Product
	}
	return service.Name
}

func (h HostItem) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s\n", h.IP))
	sb.WriteString(strings.Repeat("=", len(h.IP)) + "\n")

	for _, port := range h.Ports {
		sb.WriteString(fmt.Sprintf("%d/%s: ", port.Port, port.Proto))
		if port.Service != "" {
			sb.WriteString(fmt.Sprintf("Service: %s ", port.Service))
		}
		if port.Banner != "" {
			sb.WriteString(fmt.Sprintf("Banner: %s", port.Banner))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n")
	return sb.String()
}

