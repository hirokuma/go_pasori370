package pcd

import "errors"

// RawMsg raw data
type RawMsg struct {
	Cmd  uint8
	Data []uint8
}

// RawEncode encode RC-S956 format
func RawEncode(msg *RawMsg) []uint8 {
	data := make([]uint8, 3+2+1+1+len(msg.Data)+2)
	data[0] = 0x00
	data[1] = 0x00
	data[2] = 0xff
	data[3] = (uint8)(1 + 1 + len(msg.Data))
	data[4] = (uint8)(-data[3])
	data[5] = 0xd4
	data[6] = msg.Cmd
	copy(data[7:], msg.Data[:])

	sum := uint(0xd4 + msg.Cmd)
	for _, v := range msg.Data {
		sum += uint(v)
	}
	data[7+len(msg.Data)] = uint8(-sum)
	data[7+len(msg.Data)+1] = 0x00

	return data
}

// RawDecode decode RC-S956 format
func RawDecode(data []uint8) (*RawMsg, error) {
	if len(data) < 9 {
		// maincmd + subcmd
		return nil, errors.New("length too low")
	}
	if data[0] != 0x00 ||
		data[1] != 0x00 ||
		data[2] != 0xff {
		return nil, errors.New("preamble")
	}
	if data[len(data)-1] != 0x00 {
		return nil, errors.New("postamble")
	}
	if ((data[3] + data[4]) & 0xff) != 0x00 {
		return nil, errors.New("lcs")
	}
	sum := 0
	for idx := 5; idx < len(data); idx++ {
		sum += int(data[idx])
	}
	if (sum & 0xff) != 0x00 {
		return nil, errors.New("dcs")
	}
	// if (data[5] & 0x01) == 0 {
	// 	return nil, errors.New("not response")
	// }
	msg := new(RawMsg)
	msg.Cmd = data[6]
	msg.Data = data[7:len(data) - 2]
	return msg, nil
}
