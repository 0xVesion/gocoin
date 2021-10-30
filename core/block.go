package core

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Parent    string
	Timestamp string
	Hash      string
	Nonce     string
	Data      string
}

func (block *Block) CalcHash() string {
	h := sha256.New()
	h.Write([]byte(block.Parent + "\n" +
		block.Timestamp + "\n" +
		block.Nonce + "\n" +
		block.Data,
	))

	return hex.EncodeToString(h.Sum(nil))
}

func (block *Block) Difficulty() int {
	l := 0
	for _, r := range block.Hash {
		if r == '0' {
			l++
		}
	}

	return l
}
