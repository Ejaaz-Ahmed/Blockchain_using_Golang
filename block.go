package main

import (				//following libraries are necessary for creating the blockchain.
	"bytes"				//used for munipulating byte slices.
	"crypto/sha256"		//provides the implementation of SHA-256, cryptographic hashing algorithm.(Secure Hash Algorithm)
	"strconv"			//converting strings into numeric data and vice versa.
	"time"				//handling time related operations like calculating durations etc.
)
func (block *Block) SetHash(){			//function having pointer to structure Block. This function performs....
	timestamp:=[]byte(strconv.FormatInt(block.Timestamp, 10)) 			//converting timestamp integer value into decimal string and storing in timestamp variable.Timestamp byte slices is present in structure.go file.
	headers := bytes.Join([][]byte{timestamp,block.Previoushash,block.Alldata},[]byte{}) //here we are creating a byte slices for timestamp, Previoushash and Alldata and joining then and storing in headers.
	hash:=sha256.Sum256(headers) 		//implementation of SHA-256 algorithm for hashing headers.
	block.Newhash=hash[:]				//new value is given to Newhash byte slices present in Block structure in file structure.go
}
func NewBlock(data string,preblockhash []byte)*Block{				//function named NewBlock adds the data and timestamp and hash code and return the block.
	block:=&Block{time.Now().Unix(),preblockhash,[]byte{},[]byte(data)}		//variable having the address of structure Block.
	block.SetHash()			//sethash function is called in NewBlock function.
	return block			//return block variable. Remember, it has the address of structure Block.
}
func NewGenesisblock()*Block{ 		//genesis block is the first block in the blockchain. It is present at the start in *Block
	return NewBlock("Genesis Block",[]byte{})
}
