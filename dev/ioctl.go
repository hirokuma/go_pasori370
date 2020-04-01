package dev

// IoCtl interface
type IoCtl interface {
	Open() error
	Close()
	Send(*Msg) (*Msg, error)
}
