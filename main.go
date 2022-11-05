package main

import (
	"crypto/sha256"
	"fmt"
	"time"
)

// my block type
type Block struct {
	timestamp    time.Time
	transactions []string
	prevHash     []byte
	Hash         []byte
}

func NewBlock(transactions []string, prevHash []byte) *Block {
	currentTime := time.Now()

	return &Block{
		timestamp:    currentTime,
		transactions: transactions,
		prevHash:     prevHash,
		Hash:         NewHash(currentTime, transactions, prevHash),
	}

}

func NewHash(time time.Time, transactions []string, prevHash []byte) []byte {
	input := append(prevHash, []byte(time.String())...)
	for trans := range transactions {
		input = append(input, string(rune(trans))...)
	}
	Hash := sha256.Sum256(input)
	return Hash[:]

}

func printBlockInfo(block *Block) {
	fmt.Printf("\ttime: %s\n", block.timestamp.String())
	fmt.Printf("\tprevHash: %x\n", block.prevHash)
	fmt.Printf("\tHash: %x\n", block.Hash)
	printTransactions(block)
}

func printTransactions(block *Block) {
	fmt.Println("\tTransactions:")
	for i, transaction := range block.transactions {
		fmt.Printf("\t\t%v: %q\n", i, transaction)
	}
}

func main() {
	fmt.Println("***************************Innov**************************")
	fmt.Printf("Blockchain in golang")

	genesisTransactions := []string{"Izzy sent Will 50 bitcoin", "Will sent Izzy 30 bitcoin"}
	genesisBlock := NewBlock(genesisTransactions, []byte{})
	fmt.Println("--- First Block ---")
	printBlockInfo(genesisBlock)

	block2Transactions := []string{"John sent Izzy 30 bitcoin"}
	block2 := NewBlock(block2Transactions, genesisBlock.Hash)
	fmt.Println("--- Second Block ---")
	printBlockInfo(block2)

	block3Transactions := []string{"Will sent Izzy 45 bitcoin", "Izzy sent Will 10 bitcoin"}
	block3 := NewBlock(block3Transactions, block2.Hash)
	fmt.Println("--- Third Block ---")
	printBlockInfo(block3)
}
