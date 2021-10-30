package miner

import (
	"log"
	"time"

	"github.com/0xvesion/gocoin/core"
)

func Mine(onCancel chan interface{}, block *core.Block, difficulty int, result chan string) {
	var seconds int
	var tries int

	go func() {
		for {
			select {
			case <-onCancel:
				log.Println("Cancelling stats goroutine")
				return
			case <-time.After(time.Second):
				seconds++
				log.Printf("Mining %v @ %vkh/s\n", block.Parent[:4], tries/seconds/1000)
			}
		}
	}()

	for {
		select {
		case <-onCancel:
			log.Println("Cancelling miner goroutine")
			return
		default:
			tries++
			block.Nonce = core.GenerateNonce()
			block.Hash = block.CalcHash()

			if block.Difficulty() >= difficulty {
				result <- block.Nonce
			}
		}
	}
}
