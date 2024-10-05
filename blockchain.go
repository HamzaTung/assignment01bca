package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block structure
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
	Timestamp    time.Time
}

// Blockchain structure
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		Timestamp:    time.Now(),
	}
	block.Hash = CalculateHash(block.Transaction, block.Nonce, block.PreviousHash, block.Timestamp)
	return block
}

// AddBlock adds a block to the blockchain
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transaction, nonce, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// NewBlockChain creates the first block, the genesis block
func NewBlockChain() *Blockchain {
	genesisBlock := NewBlock("Genesis Block", 0, "")
	return &Blockchain{Blocks: []*Block{genesisBlock}}
}

// ListBlocks lists all blocks in a nice format
func (bc *Blockchain) ListBlocks() {
	fmt.Println("\nBlockchain:")
	for i, block := range bc.Blocks {
		fmt.Printf("\nBlock %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.Hash)
		fmt.Printf("Timestamp: %s\n", block.Timestamp.Format(time.RFC1123))
	}
}

// ChangeBlock changes the transaction of a specific block
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(bc.Blocks) {
		bc.Blocks[index].Transaction = newTransaction
		bc.Blocks[index].Hash = CalculateHash(bc.Blocks[index].Transaction, bc.Blocks[index].Nonce, bc.Blocks[index].PreviousHash, bc.Blocks[index].Timestamp)
	}
}

// VerifyChain verifies the blockchain for integrity
func (bc *Blockchain) VerifyChain() {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		if currentBlock.PreviousHash != previousBlock.Hash {
			fmt.Printf("\nBlockchain compromised at block %d!\n", i)
			return
		}
	}
	fmt.Println("\nBlockchain is valid!")
}

// CalculateHash generates a hash for the block
func CalculateHash(transaction string, nonce int, previousHash string, timestamp time.Time) string {
	hashInput := transaction + previousHash + fmt.Sprintf("%d", nonce) + timestamp.String()
	hash := sha256.New()
	hash.Write([]byte(hashInput))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}
