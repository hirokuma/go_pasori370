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

	for cfgNum, cfg := range dev.dev.Desc.Configs {
		config, _ := dev.dev.Config(cfgNum)
		for infNum, inf := range cfg.Interfaces {
			for altNum, alt := range inf.AltSettings {
				intfc, _ := config.Interface(infNum, altNum)
				for endNum, endPnt := range alt.Endpoints {
					if endPnt.Direction == gousb.EndpointDirectionIn {
						fmt.Printf("InEndpoint=%v\n", endPnt)
						dev.inEndpoint, _ = intfc.InEndpoint(int(endNum))
					} else {
						fmt.Printf("OutEndpoint=%v\n", endPnt)
						dev.outEndpoint, _ = intfc.OutEndpoint(int(endNum))
					}
				}
			}
		}
	}
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
