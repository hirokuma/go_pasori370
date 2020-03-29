package dev

// CmdRfConfig RFConfiguration
type CmdRfConfig struct {
	Data []uint8
}

// rfConfig RFConfiguration
func (data *CmdRfConfig) rfConfig(devPcd IoCtl) error {
	var msg Msg
	msg.Cmd = 0x32
	msg.Data = data.Data
	_, err := devPcd.Send(&msg)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigTimeout set timeout
func RfConfigTimeout(devPcd IoCtl) error {
	var data CmdRfConfig
	data.Data = []uint8{0x02, 0x00, 0x00, 0x00}
	err := data.rfConfig(devPcd)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigRetry set retry
func RfConfigRetry(devPcd IoCtl) error {
	var data CmdRfConfig
	data.Data = []uint8{0x05, 0x00, 0x00, 0x00}
	err := data.rfConfig(devPcd)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigWait set additional wait
func RfConfigWait(devPcd IoCtl) error {
	var data CmdRfConfig
	data.Data = []uint8{0x81, 0xb7}
	err := data.rfConfig(devPcd)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigOff RF off
func RfConfigOff(devPcd IoCtl) error {
	var data CmdRfConfig
	data.Data = []uint8{0x00}
	err := data.rfConfig(devPcd)
	if err != nil {
		return err
	}
	return nil
}
