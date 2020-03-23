package pcd

import (
	"errors"
	"time"

	"github.com/hirokuma/go_pasori370/dev"
)

// CodecReset reset
type CodecReset struct {
	Type int
}

// Reset reset
func (data *CodecReset) Reset() error {
	var msg dev.Msg
	msg.Cmd = 0x18
	msg.Data = []uint8{uint8(data.Type)}
	ret, err := devPcd.Send(&msg)
	if err == nil {
		if ret.Cmd == 0x19 {
			time.Sleep(time.Millisecond * 500)
		} else {
			return errors.New("")
		}
	} else {
		return errors.New("")
	}
	return nil
}
