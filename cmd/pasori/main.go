package main

import (
	"github.com/hirokuma/go_pasori370/dev"
	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	var pasori dev.Pasori370Data
	var msg pcd.RawMsg

	pasori.Open()
	defer pasori.Close()

	msg.Cmd = 0x18
	msg.Data = []uint8{0x01}
	data := pcd.RawEncode(&msg)
	pasori.Write(data)
}
