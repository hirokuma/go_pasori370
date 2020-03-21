package dev

import (
	"encoding/hex"
	"errors"
	"log"

	// https://godoc.org/github.com/google/gousb
	"github.com/google/gousb"
)

const (
	vendorID  = 0x054c
	productID = 0x02e1
)

// Pasori370Data data
type Pasori370Data struct {
	dev         *gousb.Device
	outEndpoint *gousb.OutEndpoint
	inEndpoint  *gousb.InEndpoint
}

// Open open
// https://github.com/google/gousb/blob/master/example_test.go
func (dev *Pasori370Data) Open() {
	var err error
	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open any device with a given VID/PID using a convenience function.
	dev.dev, err = ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	dev.dev.SetAutoDetach(true)

	for cfgNum, cfg := range dev.dev.Desc.Configs {
		// log.Printf("config[%d]= %v\n", cfgNum, cfg)
		for infNum, inf := range cfg.Interfaces {
			// log.Printf("interface[%d]= %v\n", infNum, inf)
			for altNum, alt := range inf.AltSettings {
				// log.Printf("alt[%d]= %v\n", altNum, alt)
				for epntNum, epnt := range alt.Endpoints {
					// log.Printf("epnt[%d]= %v\n", epntNum, epnt)
					config, err := dev.dev.Config(cfgNum)
					if err != nil {
						log.Fatalf("fail get conf: %v", err)
					}
					iface, err := config.Interface(infNum, altNum)
					if err != nil {
						log.Fatalf("fail get iface: %v", err)
					}
					// log.Printf("iface=%v\n", iface)
					err = nil
					if epnt.Direction == gousb.EndpointDirectionIn {
						dev.inEndpoint, err = iface.InEndpoint(int(epntNum))
					} else {
						dev.outEndpoint, err = iface.OutEndpoint(int(epntNum))
					}
					if err != nil {
						log.Fatalf("fail get endpoint: %v", err)
					}
				}
			}
		}
	}
	log.Printf("InEndpoint= %v\n", dev.inEndpoint)
	log.Printf("OutEndpoint= %v\n", dev.outEndpoint)
}

// Close close
func (dev *Pasori370Data) Close() {
	dev.dev.Close()
}

// Send send
func (dev *Pasori370Data) Send(msg *Msg) (*Msg, error) {
	err := dev.write(msg)
	if err != nil {
		log.Fatalf("fail write: %v\n", err)
		return nil, err
	}
	log.Printf("write\n")
	result, err := dev.read()
	if err != nil {
		log.Fatalf("fail read: %v\n", err)
		return nil, err
	}
	return result, nil
}

func (dev *Pasori370Data) write(msg *Msg) error {
	data := rawEncode(msg)
	log.Printf("data= %s\n", hex.EncodeToString(data))
	length, err := dev.outEndpoint.Write(data)
	log.Printf("wrote\n")
	if err != nil {
		return err
	}
	if length != len(data) {
		return errors.New("bad length")
	}
	return nil
}

func (dev *Pasori370Data) read() (*Msg, error) {
	preAmble := dev.rawRead(5)
	if (preAmble[0] != 0x00) || (preAmble[1] != 0x00) || (preAmble[2] != 0xff) {
		return nil, errors.New("bad preamble")
	}
	pktLen := dev.rawRead(2)
	if ((pktLen[0] + pktLen[1]) & 0xff) != 0x00 {
		return nil, errors.New("bad length")
	}
	pktData := dev.rawRead(int(pktLen[0] + 1))
	sum := 0
	for val := range pktData {
		sum += val
	}
	if (sum & 0xff) != 0x00 {
		return nil, errors.New("bad data")
	}
	postAmble := dev.rawRead(1)
	if postAmble[0] != 0x00 {
		return nil, errors.New("bad postamble")
	}
	if pktData[0] != 0xd5 {
		return nil, errors.New("not Transmit response")
	}
	result := new(Msg)
	result.Cmd = pktData[1]
	result.Data = pktData[2 : len(pktData)-1]
	return result, nil
}

func (dev *Pasori370Data) rawRead(length int) []uint8 {
	buf := make([]byte, length)
	length, err := dev.inEndpoint.Read(buf)
	if err == nil {
		buf = buf[:length]
	} else {
		buf = nil
	}
	return buf
}

// rawEncode encode RC-S956 format
func rawEncode(msg *Msg) []uint8 {
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

// rawDecode decode RC-S956 format
func rawDecode(data []uint8) (*Msg, error) {
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
	msg := new(Msg)
	msg.Cmd = data[6]
	msg.Data = data[7 : len(data)-2]
	return msg, nil
}
