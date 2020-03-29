package dev

import (
	"bytes"
	"testing"
)

func TestRawEncode_Reset0(t *testing.T) {
	var msg Msg

	msg.Cmd = 0x18
	msg.Data = []uint8{0x00}
	data := msgEncode(&msg)

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
	data := msgEncode(&msg)

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

	msg, err := msgDecode(model)
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

	msg, err := msgDecode(model)
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

func TestRawDecode_longer(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0x00, 0x11,
	}

	msg, err := msgDecode(model)
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

func TestRawDecode_invalid_length(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12,
	}

	_, err := msgDecode(model)
	if err != ErrPktInvLen {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_preamble1(t *testing.T) {
	model := []uint8{
		0x01, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadPreamble {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_preamble2(t *testing.T) {
	model := []uint8{
		0x00, 0x01, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadPreamble {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_preamble3(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0x00, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadPreamble {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_bad_lcs(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFD, 0xD5, 0x19, 0x12, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadLcs {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_bad_postamble(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x12, 0xFF,
	}

	_, err := msgDecode(model)
	if err != ErrBadPostamble {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_bad_dcs(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x19, 0x13, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadDcs {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_bad_response1(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD6, 0x19, 0x11, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadResponse {
		t.Fatalf("err: %v\n", err)
	}
}

func TestRawDecode_bad_response2(t *testing.T) {
	model := []uint8{
		0x00, 0x00, 0xFF, 0x02, 0xFE, 0xD5, 0x18, 0x13, 0x00,
	}

	_, err := msgDecode(model)
	if err != ErrBadResponse {
		t.Fatalf("err: %v\n", err)
	}
}
