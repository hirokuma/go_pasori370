package pcd

import (
	"log"

	"github.com/hirokuma/go_pasori370/dev"
)

var devPcd dev.IoCtl

// Init dev.IoCtl
func Init(device *dev.IoCtl) {
	devPcd = *device

	var dataReset MsgReset
	dataReset.Type = 0x01
	err := dataReset.Reset()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = rfConfigTimeout()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = rfConfigRetry()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = rfConfigWait()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}
