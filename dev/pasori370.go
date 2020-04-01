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

// pasori370Data data
type pasoriData struct {
	ctx         *gousb.Context
	dev         *gousb.Device
	intf        *gousb.Interface
	outEndpoint *gousb.OutEndpoint
	inEndpoint  *gousb.InEndpoint
	cfgNum      int
	infNum      int
	altNum      int
}

var pasori pasoriData

// Open open
// https://github.com/google/gousb/blob/master/example_test.go
func Open() error {
	if pasori.ctx != nil {
		return ErrAlreadyOpen
	}

	var err error
	// Initialize a new Context.
	pasori.ctx = gousb.NewContext()

	// Open any device with a given VID/PID using a convenience function.
	pasori.dev, err = pasori.ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
		return err
	}
	pasori.dev.SetAutoDetach(true)

	for cfgNum, cfg := range pasori.dev.Desc.Configs {
		// log.Printf("config[%d]= %v\n", cfgNum, cfg)
		for infNum, inf := range cfg.Interfaces {
			// log.Printf("interface[%d]= %v\n", infNum, inf)
			for altNum, alt := range inf.AltSettings {
				// log.Printf("alt[%d]= %v\n", altNum, alt)
				for epntNum, epnt := range alt.Endpoints {
					// log.Printf("epnt[%d]= %v\n", epntNum, epnt)
					config, err := pasori.dev.Config(cfgNum)
					if err != nil {
						log.Fatalf("fail get conf: %v", err)
						return err
					}
					iface, err := config.Interface(infNum, altNum)
					if err != nil {
						log.Fatalf("fail get iface: %v", err)
						return err
					}
					// log.Printf("iface=%v\n", iface)
					err = nil
					if epnt.Direction == gousb.EndpointDirectionIn {
						pasori.inEndpoint, err = iface.InEndpoint(int(epntNum))
					} else {
						pasori.outEndpoint, err = iface.OutEndpoint(int(epntNum))
					}
					if err != nil {
						log.Fatalf("fail get endpoint: %v", err)
						return err
					}
					if (pasori.inEndpoint != nil) && (pasori.outEndpoint != nil) {
						pasori.intf = iface
						pasori.cfgNum = cfgNum
						pasori.infNum = infNum
						pasori.altNum = altNum
						break
					}
				}
			}
		}
	}
	// log.Printf("iface=%v\n", pasori.intf)
	// log.Printf("InEndpoint= %v\n", pasori.inEndpoint)
	// log.Printf("OutEndpoint= %v\n", pasori.outEndpoint)

	_, err = Send(nil)
	log.Printf("Opened\n")

	return err
}

// Close close
func Close() {
	if pasori.ctx != nil {
		pasori.dev.Close()
		pasori.ctx.Close()
		pasori.ctx = nil
	}
	log.Printf("Closed\n")
}

// Send send
func Send(msg *Msg) (*Msg, error) {
	if msg == nil {
		_, err := pasori.outEndpoint.Write(ackBytes)
		return nil, err
	}
	err := write(msg)
	if err != nil {
		log.Fatalf("fail write: %v\n", err)
		return nil, err
	}
	result, err := read()
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

func write(msg *Msg) error {
	data := msgEncode(msg)
	log.Printf("write= %s\n", hex.EncodeToString(data))
	length, err := pasori.outEndpoint.Write(data)
	if err != nil {
		return err
	}
	if length != len(data) {
		return ErrSendBadLength
	}
	return nil
}

func read() (*Msg, error) {
	pkt := rawRead()
	//log.Printf("pkt1= %s\n", hex.EncodeToString(pkt))
	if bytes.Compare(ackBytes, pkt[:len(ackBytes)]) != 0 {
		return nil, ErrNotAck
	}
	if len(pkt) == len(ackBytes) {
		pkt = rawRead()
	} else {
		pkt = pkt[len(ackBytes):]
	}
	return msgDecode(pkt)
}

func rawRead() []uint8 {
	buf := make([]byte, 10*pasori.inEndpoint.Desc.MaxPacketSize)
	length, err := pasori.inEndpoint.Read(buf)
	log.Printf("read done: %d(%s)\n", length, hex.EncodeToString(buf[:length]))
	if err == nil {
		buf = buf[:length]
	} else {
		buf = nil
	}
	return buf
}
