package core

import (
	"time"
	"bytes"
	"encoding/gob"
	"log"
)

// Block keeps block headers
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result) // 创建基于result内存的编码器，别忘了&

	err := encoder.Encode(b) // 编码器对b编码,编码后的数据写入&result中
	if err != nil {
		log.Panic(err)
	}
	return result.Bytes()

}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d)) // 使用d里面的数据创建初始化reader，同时创建解码器
	err := decoder.Decode(&block) // 对d内容解码，并将解码后的数据写入&Block的内存中，别忘了&
	if err != nil {
		log.Panic(err)
	}

	return &block
}
