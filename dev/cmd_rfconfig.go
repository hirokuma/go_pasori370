package dev

// CmdRfConfig RFConfiguration
type CmdRfConfig struct {
	Data []uint8
}

// rfConfig RFConfiguration
func (data *CmdRfConfig) rfConfig() error {
	var msg Msg
	msg.Cmd = 0x32
	msg.Data = data.Data
	_, err := Send(&msg)
	if err != nil {
		return err
	}
	return nil
}

// RfConfigTimeout set timeout
func RfConfigTimeout() error {
	var data CmdRfConfig
	data.Data = []uint8{0x02, 0x00, 0x00, 0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigRetry set retry
func RfConfigRetry() error {
	var data CmdRfConfig
	data.Data = []uint8{0x05, 0x00, 0x00, 0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigWait set additional wait
func RfConfigWait() error {
	var data CmdRfConfig
	data.Data = []uint8{0x81, 0xb7}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}

// RfConfigOff RF off
func RfConfigOff() error {
	var data CmdRfConfig
	data.Data = []uint8{0x00}
	err := data.rfConfig()
	if err != nil {
		return err
	}
	return nil
}
