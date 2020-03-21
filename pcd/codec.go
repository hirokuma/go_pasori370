package pcd

import (
	"github.com/hirokuma/go_pasori370/dev"
)

// CodecReset reset
type CodecReset struct {
	Type int
}

var devPcd dev.IoCtl

// SetDevice dev.IoCtl
func SetDevice(device *dev.IoCtl) {
	devPcd = *device
}

// Reset reset
func (data *CodecReset) Reset() int {
	var msg dev.Msg
	msg.Cmd = 0x18
	msg.Data = []uint8{uint8(data.Type)}
	result := 0
	ret, err := devPcd.Send(&msg)
	if err == nil {
		if ret.Cmd != 0x19 {
			result = 1
		}
	} else {
		result = 1
	}
	return result
}
