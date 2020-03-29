package main

import (
	"log"

	"github.com/hirokuma/go_pasori370/pcd"
)

func main() {
	pcd.Open()
	log.Printf("done.\n")
	pcd.Close()
}
