package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	helpFlag    bool
	verboseFlag bool
	devFlag     string
)

func init() {
	flag.BoolVar(&helpFlag, "help", false, "Display help message")
	flag.BoolVar(&verboseFlag, "verbose", false, "Enable verbose logging")
	flag.StringVar(&devFlag, "usbdevice", "/dev/ttyUSB0", "Choose the capture device (optional)")
}

func main() {
	flag.Parse()

	if helpFlag {
		fmt.Println("This is a simple CLI application. Here are the available flags:")
		fmt.Println("-help: Display this help message")
		fmt.Println("-verbose: Enable verbose logging")
		fmt.Println("-dev: Choose the capture device (optional)")
		os.Exit(0)
	}

	if verboseFlag {
		fmt.Println("Verbose logging enabled")
	}

	if devFlag != "" {
		fmt.Printf("Capture device chosen: %s\n", devFlag)
	}

	0x100126eed
	ft, err := ftdi.OpenFirst(0x0403, 0x6001, ftdi.ChannelAny)

}
