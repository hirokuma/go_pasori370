package dev

// CmdInListPassiveTarget InListPassiveTarget
type CmdInListPassiveTarget struct {
	Data []uint8
}

// CmdInListNfcA NFC-A
type CmdInListNfcA struct {
	SensRes uint16
	SelRes  uint8
	UID     []uint8
}

// CmdInListNfcF NFC-F
type CmdInListNfcF struct {
	IDm        [8]uint8
	PMm        [8]uint8
	SystemCode uint16
}

// InListPassiveTarget InListPassiveTarget
func (data *CmdInListPassiveTarget) InListPassiveTarget() (*Msg, error) {
	var msg Msg
	msg.Cmd = 0x4a
	msg.Data = data.Data
	result, err := Send(&msg)
	return result, err
}

// InListPassiveTarget InListPassiveTarget for NFC-A
func (data *CmdInListNfcA) InListPassiveTarget() error {
	var inlist CmdInListPassiveTarget
	inlist.Data = []uint8{
		0x01, // MaxTg
		0x00, // BrTy
	}
	result, err := inlist.InListPassiveTarget()
	if err == nil {
		if result.Data[0] != 0x00 { // NbTg
			if (len(result.Data) < 6) || (result.Data[5] < 4) {
				return ErrPktInvLen
			}
			data.SensRes = uint16(result.Data[2])*16 + uint16(result.Data[3])
			data.SelRes = result.Data[4]
			data.UID = make([]uint8, result.Data[5])
			copy(data.UID, result.Data[6:6+result.Data[5]])
		}
	}
	return err
}

// InListPassiveTarget InListPassiveTarget for NFC-F
func (data *CmdInListNfcF) InListPassiveTarget() error {
	var inlist CmdInListPassiveTarget
	inlist.Data = []uint8{
		0x01, // MaxTg
		0x01, // 0x01:212Kbps  0x02:424Kbps
		0x00,
		0xff, 0xff, // SystemCode
		0x01, // opt:0x01...get SystemCode
		0x01, // Time Slot
	}
	result, err := inlist.InListPassiveTarget()
	if err == nil {
		if result.Data[0] != 0x00 {
			if (len(result.Data) < 3) || (result.Data[2] < 0x12) {
				return ErrPktInvLen
			}
			copy(data.IDm[:], result.Data[4:12])
			copy(data.PMm[:], result.Data[12:20])
			if result.Data[2] >= 0x14 {
				data.SystemCode = uint16(result.Data[20])*16 + uint16(result.Data[21])
			}
		} else {
			err = ErrNotTagFound
		}
	}
	return err
}
