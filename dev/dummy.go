package dev

import (
	"encoding/hex"
	"errors"
	"log"
)

// DummyData none
type DummyData struct {
	IsOpen bool
}

// Open open
func (dev *DummyData) Open() {
	if dev.IsOpen {
		log.Fatalf("already opened\n")
	}
	dev.IsOpen = true
	log.Printf("Opened\n")
}

// Close close
func (dev *DummyData) Close() {
	dev.IsOpen = false
	log.Printf("Closed\n")
}

// Send send
func (dev *DummyData) Send(msg *Msg) (*Msg, error) {
	if !dev.IsOpen {
		return nil, errors.New("not opened")
	}
	log.Printf("Command: 0x%02x\n", msg.Cmd)
	log.Printf("Data: %s\n", hex.EncodeToString(msg.Data))
	return nil, nil
}
