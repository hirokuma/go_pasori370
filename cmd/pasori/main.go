package main

import (
	"log"

	"github.com/hirokuma/go_pasori370/dev"
	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	var device dev.Pasori370Data

	device.Open()
	defer device.Close()
	var devdev dev.IoCtl = &device
	pcd.Init(&devdev)

	log.Printf("done.\n")
}
