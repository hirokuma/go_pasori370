package dev

// IoCtl interface
type IoCtl interface {
	Open()
	Close()
	Send(*Msg) (*Msg, error)
}
