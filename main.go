package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var predefinedProfiles = map[string][]int{
	"S7-300":       {102, 502},
	"S7-1200":      {102, 502, 8080},
	"S7-1500":      {102, 502, 161, 443},
	"ScalanceX200": {80, 443},
	"IPC":          {3389}, // Add the RDP port (3389) for IPC devices
	"SiemensHMI":   {80, 102, 161}, // Add the specific ports for Siemens HMI devices
	"All": 			{102, 502, 8080, 161, 443, 80, 3389}, // Combine all predefined profiles
	"Custom":       {},
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

	fmt.Print("Select a predefined profile (S7-300, S7-1200, S7-1500, ScalanceX200, IPC, SiemensHMI, All, Custom): ")
	fmt.Scan(&input)
	selectedProfile, profileExists := predefinedProfiles[input]

	if !profileExists {
		log.Fatal("Invalid profile. Use 'S7-300', 'S7-1200', 'S7-1500', 'ScalanceX200', 'IPC', 'SiemensHMI', 'All', or 'Custom'.")
	}

	// Generate a unique log file name with the current date and time
	currentDateTime := time.Now().Format("2006-01-02-15-04-05")
	logFileName := fmt.Sprintf("scan_results_%s.txt", currentDateTime)

	logFile, err := os.Create(logFileName)
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

			// Read and log the banner information
			banner := make([]byte, 1024)
			_, err = conn.Read(banner)
			if err != nil {
				banner = []byte("No banner information available.")
			}

			conn.Close()
			result := fmt.Sprintf("Port %d is open on %s\nBanner Information: %s\n", port, ip, string(banner))
			log.Print(result)
		}
	}
}
