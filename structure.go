package main
type Block struct{			//structure named Block having....
	Timestamp int64			//Timestamp int type.
	Previoushash []byte		//Previoushash Byte slice.
	Newhash []byte			//Newhash Byte slice.
	Alldata []byte			//Alldata Byte slice.
}

type Blockchain struct {	//Structure named Blockchain. 
	Blocks []*Block			//Blocks byte slice pointer to Block structure.
}