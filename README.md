# GO-OT-scanning

This is a simple command-line port scanning tool written in Golang. It allows you to scan a range of IP addresses for specific port numbers and logs the results to a text file with the current date and time.

## Usage

1. Clone the repository or download the `main.go` file.

2. Ensure you have Go (Golang) installed on your system.

3. Run the tool by executing the following command in the terminal: "go run main.go"

4. Follow the on-screen prompts:

- Enter the IP address range to scan in the format 'start-end' (e.g., "192.168.1.1-254").

- Select a predefined profile for scanning. Available profiles include:
  - S7-300
  - S7-1200
  - S7-1500
  - ScalanceX200
  - IPC (RDP)
  - SiemensHMI
  - All (Scans all predefined profiles except Custom)
  - Custom (Specify custom port numbers to scan)

1. The tool will scan the specified IP address range for the selected profile's port numbers and log the results in a text file with a unique name that includes the current date and time (e.g., "scan_results_2023-10-13-15-30-45.txt").

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### MIT License

MIT License

Copyright (c) 2023 Konstantinos Poumpouridis

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS," WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE, AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES, OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT, OR OTHERWISE, ARISING FROM, OUT OF, OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.