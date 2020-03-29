package dev

import (
	"bytes"
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

var ackBytes = []uint8{0x00, 0x00, 0xff, 0x00, 0xff, 0x00}

// Pasori370Data data
type Pasori370Data struct {
	ctx         *gousb.Context
	dev         *gousb.Device
	intf        *gousb.Interface
	outEndpoint *gousb.OutEndpoint
	inEndpoint  *gousb.InEndpoint
	cfgNum      int
	infNum      int
	altNum      int
}

// Open open
// https://github.com/google/gousb/blob/master/example_test.go
func (dev *Pasori370Data) Open() {
	var err error
	// Initialize a new Context.
	dev.ctx = gousb.NewContext()

	// Open any device with a given VID/PID using a convenience function.
	dev.dev, err = dev.ctx.OpenDeviceWithVIDPID(vendorID, productID)
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
					if (dev.inEndpoint != nil) && (dev.outEndpoint != nil) {
						dev.intf = iface
						dev.cfgNum = cfgNum
						dev.infNum = infNum
						dev.altNum = altNum
						break
					}
				}
			}
		}
	}
	log.Printf("iface=%v\n", dev.intf)
	log.Printf("InEndpoint= %v\n", dev.inEndpoint)
	log.Printf("OutEndpoint= %v\n", dev.outEndpoint)

	dev.Send(nil)
}

// Close close
func (dev *Pasori370Data) Close() {
	dev.dev.Close()
	dev.ctx.Close()
}

// Send send
func (dev *Pasori370Data) Send(msg *Msg) (*Msg, error) {
	if msg == nil {
		_, err := dev.outEndpoint.Write(ackBytes)
		return nil, err
	}
	err := dev.write(msg)
	if err != nil {
		log.Fatalf("fail write: %v\n", err)
		return nil, err
	}
	result, err := dev.read()
	if err != nil {
		log.Fatalf("fail read: %v\n", err)
		return nil, err
	}
	if result.Cmd != msg.Cmd+1 {
		err = errors.New("bad sub response")
		log.Fatalf("fail: %v\n", err)
		return nil, err
	}
	return result, nil
}

func (dev *Pasori370Data) write(msg *Msg) error {
	data := msgEncode(msg)
	log.Printf("write= %s\n", hex.EncodeToString(data))
	length, err := dev.outEndpoint.Write(data)
	if err != nil {
		return err
	}
	if length != len(data) {
		return ErrSendBadLength
	}
	return nil
}

func (dev *Pasori370Data) read() (*Msg, error) {
	pkt := dev.rawRead()
	//log.Printf("pkt1= %s\n", hex.EncodeToString(pkt))
	if bytes.Compare(ackBytes, pkt[:len(ackBytes)]) != 0 {
		return nil, ErrNotAck
	}
	if len(pkt) == len(ackBytes) {
		pkt = dev.rawRead()
	} else {
		pkt = pkt[len(ackBytes):]
	}
	return msgDecode(pkt)
}

func (dev *Pasori370Data) rawRead() []uint8 {
	buf := make([]byte, 10*dev.inEndpoint.Desc.MaxPacketSize)
	length, err := dev.inEndpoint.Read(buf)
	log.Printf("read done: %d(%s)\n", length, hex.EncodeToString(buf[:length]))
	if err == nil {
		buf = buf[:length]
	} else {
		buf = nil
	}
	return buf
}
