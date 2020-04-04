package main

import (
	"encoding/hex"
	"log"

	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	pcd.Open()

	var poll pcd.ResultPolling
	err := poll.Polling()
	if err == nil {
		log.Printf("IDm: %s\n", hex.EncodeToString(poll.NfcF.IDm[:]))
		log.Printf("PMm: %s\n", hex.EncodeToString(poll.NfcF.PMm[:]))
		log.Printf("SystemCode: 0x%04x\n", poll.NfcF.SystemCode)
	}

	pcd.Close()
}
