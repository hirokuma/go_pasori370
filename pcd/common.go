package pcd

import "github.com/hirokuma/go_pasori370/dev"

var devPcd dev.IoCtl

// SetDevice dev.IoCtl
func SetDevice(device *dev.IoCtl) {
	devPcd = *device
}
