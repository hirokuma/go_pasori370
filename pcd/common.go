package pcd

import (
	"log"

	"github.com/hirokuma/go_pasori370/dev"
)

// ResultPolling result
type ResultPolling struct {
	NfcA *dev.CmdInListNfcA
	NfcF *dev.CmdInListNfcF
}

// Open open Proximity Coupling Device
func Open() {
	err := dev.Open()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}

	var dataReset dev.CmdReset
	dataReset.Type = 0x01
	err = dataReset.Reset()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigTimeout()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigRetry()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	err = dev.RfConfigWait()
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
}

// Close close
func Close() {
	dev.Close()
}

// Polling polling
func (result *ResultPolling) Polling() error {
	{
		var inlist dev.CmdInListNfcF
		err := inlist.InListPassiveTarget()
		if err == nil {
			result.NfcF = &inlist
			return nil
		}
	}

	{
		var inlist dev.CmdInListNfcA
		err := inlist.InListPassiveTarget()
		if err == nil {
			result.NfcA = &inlist
			return nil
		}
	}
	return dev.ErrNotTagFound
}
