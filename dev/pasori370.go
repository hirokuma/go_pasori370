package dev

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

const (
	vendorID  = 0x054c
	productID = 0x02e1
)

// Pasori370Data data
type Pasori370Data struct {
	dev         *gousb.Device
	outEndpoint *gousb.OutEndpoint
	inEndpoint  *gousb.InEndpoint
}

// Open open
// https://github.com/google/gousb/blob/master/example_test.go
func (dev *Pasori370Data) Open() {
	var err error
	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	dev.dev, err = ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	dev.dev.SetAutoDetach(true)

	config, err := dev.dev.Config(1)
	if err != nil {
		log.Fatalf("fail get config: %v", err)
	}
	fmt.Printf("config= %v\n", config)
	intr, err := config.Interface(0, 0)
	if err != nil {
		log.Fatalf("fail get interface: %v", err)
	}
	fmt.Printf("interface= %v\n", intr)
	dev.inEndpoint, err = intr.InEndpoint(4)
	if err != nil {
		log.Fatalf("fail get InEndpoint: %v", err)
	}
	dev.outEndpoint, err = intr.OutEndpoint(4)
	if err != nil {
		log.Fatalf("fail get OutEndpoint: %v", err)
	}
	fmt.Printf("InEndpoint= %v\n", dev.inEndpoint)
	fmt.Printf("OutEndpoint= %v\n", dev.outEndpoint)
	fmt.Printf("\n\n")

	for cfgNum, cfg := range dev.dev.Desc.Configs {
		fmt.Printf("config[%d]= %v\n", cfgNum, cfg)
		for infNum, inf := range cfg.Interfaces {
			fmt.Printf("interface[%d]= %v\n", infNum, inf)
			for altNum, alt := range inf.AltSettings {
				fmt.Printf("alt[%d]= %v\n", altNum, alt)
				for epntNum, epnt := range alt.Endpoints {
					fmt.Printf("epnt[%d]= %v\n", epntNum, epnt)
					config, err = dev.dev.Config(cfgNum)
					if err != nil {
						log.Fatalf("fail get conf!!: %v", err)
					}
					iface, err := config.Interface(infNum, altNum)
					if err != nil {
						log.Fatalf("fail get iface!!: %v", err)
					}
					fmt.Printf("iface=%v\n", iface)
					if epnt.Direction == gousb.EndpointDirectionIn {
						dev.inEndpoint, err = iface.InEndpoint(int(epntNum))
					} else {
						dev.outEndpoint, err = iface.OutEndpoint(int(epntNum))
					}
				}
			}
		}
	}
	fmt.Printf("InEndpoint= %v\n", dev.inEndpoint)
	fmt.Printf("OutEndpoint= %v\n", dev.outEndpoint)
	fmt.Printf("\n\n")
}

// Close close
func (dev *Pasori370Data) Close() {
	dev.dev.Close()
}

// Read read
func (dev *Pasori370Data) Read(int) []uint8 {
	return nil
}

// Write write
func (dev *Pasori370Data) Write(data []uint8) bool {
	return true
}
