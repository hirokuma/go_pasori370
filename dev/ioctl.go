package dev

// IoCtl interface
type IoCtl interface {
	Open() bool
	Close()
	Read(int) []uint8
	Write([]uint8) bool
}
