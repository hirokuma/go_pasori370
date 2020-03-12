package dev

import (
	"encoding/hex"
	"log"
)

// DummyData none
type DummyData struct {
	IsOpen bool
}

// Open open
func (dev *DummyData) Open() bool {
	if dev.IsOpen {
		log.Fatalf("already opened\n")
	}
	dev.IsOpen = true
	log.Printf("Opened\n")
	return true
}

// Close close
func (dev *DummyData) Close() {
	dev.IsOpen = false
	log.Printf("Closed\n")
}

// Read read
func (dev *DummyData) Read(int) []uint8 {
	if !dev.IsOpen {
		log.Printf("not opened\n")
	}
	log.Printf("Read\n")
	return nil
}

// Write write
func (dev *DummyData) Write(data []uint8) bool {
	if !dev.IsOpen {
		log.Printf("not opened\n")
		return false
	}
	log.Printf("WriteData: %s\n", hex.Dump(data))
	return true
}
