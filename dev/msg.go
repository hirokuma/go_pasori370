package dev

// Msg raw data
type Msg struct {
	Cmd  uint8
	Data []uint8
}

// msgEncode encode RC-S956 format
func msgEncode(msg *Msg) []uint8 {
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

// msgDecode decode RC-S956 format
func msgDecode(pkt []uint8) (*Msg, error) {
	//log.Printf("pkt2= %s\n", hex.EncodeToString(pkt))
	if len(pkt) < 9 {
		// maincmd + subcmd
		return nil, ErrPktInvLen
	}
	if (pkt[0] != 0x00) || (pkt[1] != 0x00) || (pkt[2] != 0xff) {
		return nil, ErrBadPreamble
	}
	if ((pkt[3] + pkt[4]) & 0xff) != 0x00 {
		return nil, ErrBadLcs
	}

	cmd := pkt[5:]
	datalen := pkt[3]
	// (datalen) + DCS + 00
	if cmd[datalen+1] != 0x00 {
		return nil, ErrBadPostamble
	}
	sum := 0
	for _, val := range cmd[:datalen+1] {
		sum += int(val)
	}
	if (sum & 0xff) != 0x00 {
		return nil, ErrBadDcs
	}
	if cmd[0] != 0xd5 {
		return nil, ErrBadResponse
	}
	if (cmd[1] % 2) != 1 {
		return nil, ErrBadResponse
	}
	result := new(Msg)
	result.Cmd = cmd[1]
	if datalen > 2 {
		result.Data = cmd[2:datalen]
	}
	//log.Printf("cmd=%02x, data=%s\n", result.Cmd, hex.EncodeToString(result.Data))
	return result, nil
}
