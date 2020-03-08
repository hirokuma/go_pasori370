package pcd_test

import (
	"bytes"
	"testing"

	"github.com/hirokuma/go_pasori370/pcd"
)

func TestRawEncode_Reset0(t *testing.T) {
	var msg pcd.RawMsg

	msg.Cmd = 0x18
	msg.Data = []uint8{0x00}
	data := pcd.RawEncode(&msg)

	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x00, 0x14, 0x00,
	}
	if bytes.Compare(model, data) != 0 {
		t.Fatalf("not same\n")
	}
}

func TestRawEncode_Reset1(t *testing.T) {
	var msg pcd.RawMsg

	msg.Cmd = 0x18
	msg.Data = []uint8{0x01}
	data := pcd.RawEncode(&msg)

	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x01, 0x13, 0x00,
	}
	if bytes.Compare(model, data) != 0 {
		t.Fatalf("not same\n")
	}
}

func TestRawDecode_Reset0(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x00, 0x14, 0x00,
	}

	msg, err := pcd.RawDecode(model)
	if err != nil {
		t.Fatalf("err\n")
	}
	if msg.Cmd != 0x18 {
		t.Fatalf("not 0x18\n")
	}
	if len(msg.Data) != 1 {
		t.Fatalf("invalid length\n")
	}
	if msg.Data[0] != 0x00 {
		t.Fatalf("invalid data\n")
	}
}

func TestRawDecode_Reset1(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x03, 0xFD, 0xD4, 0x18, 0x01, 0x13, 0x00,
	}

	msg, err := pcd.RawDecode(model)
	if err != nil {
		t.Fatalf("err\n")
	}
	if msg.Cmd != 0x18 {
		t.Fatalf("not 0x18\n")
	}
	if len(msg.Data) != 1 {
		t.Fatalf("invalid length\n")
	}
	if msg.Data[0] != 0x01 {
		t.Fatalf("invalid data\n")
	}
}
