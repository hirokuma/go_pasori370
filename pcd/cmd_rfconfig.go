package pcd

import (
	"github.com/hirokuma/go_pasori370/dev"
)

// MsgRfConfig RFConfiguration
type MsgRfConfig struct {
	Data []uint8
}

// RfConfig RFConfiguration
func (data *MsgRfConfig) RfConfig() error {
	var msg dev.Msg
	msg.Cmd = 0x32
	msg.Data = data.Data
	_, err := devPcd.Send(&msg)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigTimeout set timeout
func RfConfigTimeout() error {
	var data MsgRfConfig
	data.Data = []uint8{0x02, 0x00, 0x00, 0x00}
	err := data.RfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigRetry set retry
func RfConfigRetry() error {
	var data MsgRfConfig
	data.Data = []uint8{0x05, 0x00, 0x00, 0x00}
	err := data.RfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigWait set additional wait
func RfConfigWait() error {
	var data MsgRfConfig
	data.Data = []uint8{0x81, 0xb7}
	err := data.RfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigOff RF off
func RfConfigOff() error {
	var data MsgRfConfig
	data.Data = []uint8{0x00}
	err := data.RfConfig()
	if err != nil {
		return err
	}
	return nil
}
