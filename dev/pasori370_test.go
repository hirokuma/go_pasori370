package dev

import (
	"bytes"
	"testing"
)

func TestRawEncode_Reset0(t *testing.T) {
	var msg Msg

	msg.Cmd = 0x18
	msg.Data = []uint8{0x00}
	data := rawEncode(&msg)

	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x00, 0x14, 0x00,
	}
	if bytes.Compare(model, data) != 0 {
		t.Fatalf("not same\n")
	}
}

func TestRawEncode_Reset1(t *testing.T) {
	var msg Msg

	msg.Cmd = 0x18
	msg.Data = []uint8{0x01}
	data := rawEncode(&msg)

	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x01, 0x13, 0x00,
	}
	if bytes.Compare(model, data) != 0 {
		t.Fatalf("not same\n")
	}
}

func TestRawDecode_nodata(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0x00,
	}

	msg, err := rawDecode(model)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if msg.Cmd != 0x19 {
		t.Fatalf("not 0x19\n")
	}
	if len(msg.Data) != 0 {
		t.Fatalf("invalid length: %d\n", len(msg.Data))
	}
}

func TestRawDecode_withdata(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD5, 0x19, 0xAB, 0x67, 0x00,
	}

	msg, err := rawDecode(model)
	if err != nil {
		t.Fatalf("err: %v\n", err)
	}
	if msg.Cmd != 0x19 {
		t.Fatalf("not 0x19\n")
	}
	if len(msg.Data) != 1 {
		t.Fatalf("invalid length: %d\n", len(msg.Data))
	}
	if msg.Data[0] != 0xab {
		t.Fatalf("bad data: %v\n", msg.Data)
	}
}

func TestRawDecode_tooshort(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12,
	}

	_, err := rawDecode(model)
	if err == nil {
		t.Fatalf("err: %v\n", err)
	}
}
