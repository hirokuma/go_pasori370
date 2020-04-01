package dev

import (
	"time"
)

// CmdReset reset
type CmdReset struct {
	Type int
}

// Reset reset
func (data *CmdReset) Reset() error {
	var msg Msg
	msg.Cmd = 0x18
	msg.Data = []uint8{uint8(data.Type)}
	_, err := Send(&msg)
	if err == nil {
		time.Sleep(time.Millisecond * 500)
	} else {
		return err
	}
	return nil
}
