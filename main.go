package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var predefinedProfiles = map[string][]int{
	"S7-300":   {102, 502},
	"S7-1200":  {102, 502, 8080},
	"S7-1500":  {102, 502, 161, 443},
	"Custom":   {},
}

func main() {
	fmt.Print("Enter the IP address range to scan (e.g., 192.168.1.1-254): ")
	var input string
	fmt.Scan(&input)

	ipRange := strings.Split(input, "-")
	if len(ipRange) != 2 {
		log.Fatal("Invalid IP address range format. Use the format 'start-end'.")
	}

	startIP := ipRange[0]
	endIP := ipRange[1]

	// Validate and convert IP addresses to integer representation
	startIPInt := ipToInt(startIP)
	endIPInt := ipToInt(endIP)

	if startIPInt == -1 || endIPInt == -1 {
		log.Fatal("Invalid IP address in the range.")
	}

	fmt.Print("Select a predefined profile (S7-300, S7-1200, S7-1500, Custom): ")
	fmt.Scan(&input)
	selectedProfile, profileExists := predefinedProfiles[input]

	if !profileExists {
		log.Fatal("Invalid profile. Use 'S7-300', 'S7-1200', 'S7-1500', or 'Custom'.")
	}

	logFile, err := os.Create("scan_results.txt")
	if err != nil {
		log.Fatal("Cannot create log file")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	for i := startIPInt; i <= endIPInt; i++ {
		ip := intToIP(i)
		for _, port := range selectedProfile {
			address := fmt.Sprintf("%s:%d", ip, port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				continue // Port is closed or filtered
			}
			conn.Close()
			result := fmt.Sprintf("Port %d is open on %s\n", port, ip)
			log.Print(result)
		}
	}
}

// Function to convert an IP address to an integer
func ipToInt(ip string) int {
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return -1
	}

	var result int
	for _, part := range parts {
		partInt, err := strconv.Atoi(part)
		if err != nil || partInt < 0 || partInt > 255 {
			return -1
		}
		result = (result << 8) | partInt
	}
	return result
}

// Function to convert an integer to an IP address
func intToIP(ipInt int) string {
	return fmt.Sprintf("%d.%d.%d.%d", (ipInt>>24)&0xFF, (ipInt>>16)&0xFF, (ipInt>>8)&0xFF, ipInt&0xFF)
}
