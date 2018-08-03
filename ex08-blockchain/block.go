package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte, bits int) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block, bits)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (bc *Blockchain) StageBlock(data string) {
	fmt.Printf("Your block with data %s was staged, now you need to mine it\n", data)
}

func (bc *Blockchain) AddBlock(data string, bits int) *Block {
	if bits < 1 {
		panic("[ERROR] You can't mine with the difficulty 0")
	}
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash, bits)
	bc.blocks = append(bc.blocks, newBlock)
	fmt.Printf("\n[NEW] block was succesfully added!\n")
	fmt.Printf("Block's data is:%s\n", data)
	return newBlock
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{}, 0)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func (b *Blockchain) ListBlocks() {
	fmt.Printf("<<< LIST OF ALL BLOCKS >>>\n")
	for _, block := range b.blocks {
		PrintBlock(block)
		fmt.Println()
	}
}

func PrintBlock(b *Block) {
	fmt.Printf("Previous block hash: %x\n", b.PrevBlockHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println()
}
