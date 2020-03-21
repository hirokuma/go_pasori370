package main

import (
	"github.com/hirokuma/go_pasori370/dev"
	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	var device dev.Pasori370Data
	//var device dev.DummyData

	device.Open()
	defer device.Close()
	var devdev dev.IoCtl = &device
	pcd.SetDevice(&devdev)

	var dataReset pcd.CodecReset
	dataReset.Type = 0x01
	dataReset.Reset()
}
