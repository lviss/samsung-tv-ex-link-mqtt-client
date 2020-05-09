package main

	import (
		"github.com/mgoff/go-samsung-exlink"
		"log"
	)

	func main() {

		// open the connection to the EX-Link device
		device, err := exlink.Open("/dev/ttyUSB0")
		if err != nil {
			log.Fatal(err)
		}

		// close the connection at the end
		defer device.Close()

		// switch to HDMI 1
		err = device.SourceHDMI1()
		if err != nil {
			log.Fatal(err)
		}
	}
