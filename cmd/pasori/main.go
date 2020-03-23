package main

import (
	"log"

	"github.com/hirokuma/go_pasori370/dev"
	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	var device dev.Pasori370Data
	//var device dev.DummyData

	device.Open()
	defer device.Close()
	var devdev dev.IoCtl = &device
	pcd.SetDevice(&devdev)

	var dataReset pcd.MsgReset
	dataReset.Type = 0x01
	err := dataReset.Reset()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	err = pcd.RfConfigTimeout()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = pcd.RfConfigRetry()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = pcd.RfConfigWait()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	log.Printf("done.\n")
}
