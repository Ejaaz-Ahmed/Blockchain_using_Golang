package main 	//using the main package.

import (		//from main package we import "fmt" to print the data on screen.
	"fmt"
)

func main() {										//starting the main function.
	fmt.Println("Creating first blockchain:")					//printing message.
	newBlockchain := NewBlockchain()						//newBlockchain is the function of type struct defined in blockchain.go file.	
	newBlockchain.Addblock("First transaction")					//using newBlockchain object, Addblock function is called which is defined in blockchain.go file.
	newBlockchain.Addblock("Second transcation")		
	for i, block := range newBlockchain.Blocks {					//for is used to print the blockchain(chaining the blocks together) and its range is number of blocks added.
		fmt.Printf("Block ID:%d\n", i)								//printing block ID.
		fmt.Printf("Timestamp: %d\n", block.Timestamp+int64(i))		//timestamp is the exact time on which the block is created and added into the blockchain.
		fmt.Printf("Hash of Block %x\n", block.Newhash)				//printing the hash of block.
		fmt.Printf("Hash of Previous Block: %x\n", block.Previoushash)		//printing the previous hash of block.
		fmt.Printf("Transactions number: %s\n", block.Alldata)			//prints the number of transaction.
	}
}
