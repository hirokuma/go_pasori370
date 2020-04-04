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
		if poll.NfcF != nil {
			log.Printf("Found: NFC-F\n")
			log.Printf("IDm: %s\n", hex.EncodeToString(poll.NfcF.IDm[:]))
			log.Printf("PMm: %s\n", hex.EncodeToString(poll.NfcF.PMm[:]))
			log.Printf("SystemCode: 0x%04x\n", poll.NfcF.SystemCode)
		} else if poll.NfcA != nil {
			log.Printf("Found: NFC-A\n")
			log.Printf("SENS_RES: 0x%04x\n", poll.NfcA.SensRes)
			log.Printf("SEL_RES: 0x%02x\n", poll.NfcA.SelRes)
			log.Printf("UID: %s\n", hex.EncodeToString(poll.NfcA.UID))
		}
	}

	pcd.Close()
}
