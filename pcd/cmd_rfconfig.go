package pcd

import (
	"github.com/hirokuma/go_pasori370/dev"
)

// MsgRfConfig RFConfiguration
type MsgRfConfig struct {
	Data []uint8
}

// rfConfig RFConfiguration
func (data *MsgRfConfig) rfConfig() error {
	var msg dev.Msg
	msg.Cmd = 0x32
	msg.Data = data.Data
	_, err := devPcd.Send(&msg)
	if err != nil {
		return err
	}
	return nil
}

// rfConfigTimeout set timeout
func rfConfigTimeout() error {
	var data MsgRfConfig
	data.Data = []uint8{0x02, 0x00, 0x00, 0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// rfConfigRetry set retry
func rfConfigRetry() error {
	var data MsgRfConfig
	data.Data = []uint8{0x05, 0x00, 0x00, 0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// rfConfigWait set additional wait
func rfConfigWait() error {
	var data MsgRfConfig
	data.Data = []uint8{0x81, 0xb7}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigOff RF off
func RfConfigOff() error {
	var data MsgRfConfig
	data.Data = []uint8{0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}
