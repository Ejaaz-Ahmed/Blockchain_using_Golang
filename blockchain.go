package main
func (blockchain *Blockchain) Addblock(data string){			//function pointer to Blockchain structure in which blocks are added in byte slice
	Previousblock:=blockchain.Blocks[len(blockchain.Blocks)-1]	//previousblock variable.
	newBlock:=NewBlock(data,Previousblock.Newhash)				//newBlock variable having data and newBlock function defined in block.go.
	blockchain.Blocks=append(blockchain.Blocks, newBlock)		//blockchain variable accessing Blocks []byte in Blockchain structure adding new blocks.
}
func NewBlockchain() * Blockchain{								//NewBlockchain function pointer to Blockchain structure returning the whole blockchain with genesis block in first.
	return &Blockchain{[]*Block{NewGenesisblock()}}
}