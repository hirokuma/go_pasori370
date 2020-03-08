package main

import (
	"github.com/hirokuma/go_pasori370/dev"
	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	var dummy dev.DummyData
	var msg pcd.RawMsg

	dummy.Open()
	defer dummy.Close()

	msg.Cmd = 0x18
	msg.Data = []uint8{0x01}
	data := pcd.RawEncode(&msg)
	dummy.Write(data)
}
