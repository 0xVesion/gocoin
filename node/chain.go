package node

import (
	"log"
	"sync"

	"github.com/0xvesion/gocoin/core"
)

type Chain struct {
	Blocks        []core.Block
	Current       core.Block
	Difficulty    int
	BlockSize     int
	EpochDuration int
	BlockDuration int
	OnUpdate      func()
	w             sync.Mutex
}

func NewChain() *Chain {
	genesis := core.Block{
		Parent:    "⛓",
		Timestamp: core.CurrentTimestamp(),
		Data:      "This is just the beginning.",
	}
	genesis.Hash = genesis.CalcHash()

	chain := Chain{
		BlockSize:     0xFF,
		BlockDuration: 0xF,
		EpochDuration: 0xF,
		Difficulty:    1,

		Blocks:   []core.Block{genesis},
		OnUpdate: nil,
		w:        sync.Mutex{},
	}

	chain.Current = chain.NewBlock()

	chain.emitUpdate()

	return &chain
}

func (chain *Chain) NewBlock() core.Block {
	return core.Block{
		Parent:    chain.LastBlock().Hash,
		Timestamp: core.CurrentTimestamp(),
	}
}
func (chain *Chain) Height() int {
	return len(chain.Blocks)
}

func (chain *Chain) LastBlock() core.Block {
	return chain.Blocks[chain.Height()-1]
}

func (chain *Chain) IsValid() bool {
	for i := 1; i < len(chain.Blocks); i++ {
		previous := chain.Blocks[i-1]
		current := chain.Blocks[i]

		if current.Parent != previous.Hash {
			return false
		}

		if current.Hash != current.CalcHash() {
			return false
		}
	}

	return true
}

func (chain *Chain) AdvanceBlock(nonce string) bool {
	chain.w.Lock()
	defer chain.w.Unlock()

	c := chain.Current
	c.Nonce = nonce
	c.Hash = c.CalcHash()

	if c.Difficulty() < chain.Difficulty {
		return false
	}

	chain.Blocks = append(chain.Blocks, c)

	if !chain.IsValid() {
		panic("Invalid chain")
	}

	chain.updateDifficulty()
	chain.commit()
	log.Printf("New Block. height: %v difficulty: %v epoch: %v/%v\n",
		chain.Height(), chain.Difficulty, chain.Height()%chain.EpochDuration, chain.EpochDuration)

	chain.Current = chain.NewBlock()
	chain.emitUpdate()

	return true
}

func (chain *Chain) AppendData(data string) bool {
	chain.w.Lock()
	defer chain.w.Unlock()

	if len(data)+len(chain.Current.Data) > chain.BlockSize {
		return false
	}

	chain.Current.Data += data
	chain.emitUpdate()

	return true
}

func (chain *Chain) updateDifficulty() {
	isStartOfNewEpoch := chain.Height()%chain.EpochDuration != 0
	if isStartOfNewEpoch {
		return
	}

	start := core.ParseTimestamp(chain.Blocks[len(chain.Blocks)-chain.EpochDuration].Timestamp)
	end := core.ParseTimestamp(chain.Current.Timestamp)

	duration := end.Sub(start).Seconds()

	if duration/float64(chain.EpochDuration) > float64(chain.BlockDuration) {
		chain.Difficulty--
	} else {
		chain.Difficulty++
	}
}

func (chain *Chain) emitUpdate() {
	go func() {
		if chain.OnUpdate == nil {
			return
		}

		chain.OnUpdate()
	}()
}

func (chain *Chain) commit() {
	// TODO: Implement
}
