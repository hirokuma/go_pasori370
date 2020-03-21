package dev

// Msg raw data
type Msg struct {
	Cmd  uint8
	Data []uint8
}

// IoCtl interface
type IoCtl interface {
	Open()
	Close()
	Send(*Msg) (*Msg, error)
}
