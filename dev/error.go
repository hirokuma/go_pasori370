package dev

import "errors"

var (
	// ErrPktInvLen invalid packet length
	ErrPktInvLen = errors.New("invalid length")

	// ErrBadPreamble bad preamble
	ErrBadPreamble = errors.New("bad preamble")

	// ErrBadLcs bad LCS
	ErrBadLcs = errors.New("bad length")

	// ErrBadDcs bad DCS
	ErrBadDcs = errors.New("bad data")

	// ErrBadPostamble bad preamble
	ErrBadPostamble = errors.New("bad postamble")

	// ErrBadResponse bad response
	ErrBadResponse = errors.New("bad response")

	// ErrSendBadLength send bad length
	ErrSendBadLength = errors.New("bad length")

	// ErrNotAck not ACK
	ErrNotAck = errors.New("not ACK")
)
