package pcd

import (
	"time"

	"github.com/hirokuma/go_pasori370/dev"
)

// MsgReset reset
type MsgReset struct {
	Type int
}

// Reset reset
func (data *MsgReset) Reset() error {
	var msg dev.Msg
	msg.Cmd = 0x18
	msg.Data = []uint8{uint8(data.Type)}
	_, err := devPcd.Send(&msg)
	if err == nil {
		time.Sleep(time.Millisecond * 500)
	} else {
		return err
	}
	return nil
}
