package pcd

import (
	"log"

	"github.com/hirokuma/go_pasori370/dev"
)

var devPcd dev.IoCtl

// Open dev.IoCtl
func Open() {
	if devPcd != nil {
		log.Fatalf("already opened\n")
	}

	var device dev.Pasori370Data

	device.Open()
	devPcd = &device

	var dataReset dev.CmdReset
	dataReset.Type = 0x01
	err := dataReset.Reset(devPcd)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigTimeout(devPcd)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigRetry(devPcd)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigWait(devPcd)
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

// Close close
func Close() {
	if devPcd != nil {
		devPcd.Close()
		devPcd = nil
	}
}
