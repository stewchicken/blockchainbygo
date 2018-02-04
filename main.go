
package main

import "bytes"
import "crypto/sha256"
import "strconv"
import "time"
import "fmt"
import "os"

//step1 declare a Block
type Block struct{
	Timestamp int64 //here to store the  created time of the Block
	Data []byte //here to store data( for example the bitcoin or other information)
	PrevBlockHash []byte // here to store the hashvalue of last Block
	Hash []byte //here to store the hashvalue of current Block
} 

//step2 hash a BlockChain with SHA256
func (this *Block) SetHash(){
	//remove milliseconds of BlockChain's timestamp
	timestamp :=[]byte(strconv.FormatInt(this.Timestamp,10))
	//concatenate the Blocckchian's PrevBlockHash, Data,Timestamp with emtpy []byte{}
	//Block's prevBlockHash is also encrypted with Hashvalue, if Block's value is modified, all Blocks in Blockchain must be modified
	//therefore it is difficult to hack.
	headers := bytes.Join([][]byte{this.PrevBlockHash,this.Data,timestamp},[]byte{})
	//hash headers with SHA256
	hash :=sha256.Sum256(headers)
	this.Hash=hash[:]
}

//step3 , create a Block
func NewBlock (data string, prevBlockHash[]byte) *Block{
	//create a Block
	block := Block{}
	block.Timestamp=time.Now().Unix()
	block.Data=[]byte(data)
	block.PrevBlockHash=prevBlockHash
	block.Hash=[]byte{}
	block.SetHash()
	return &block
}

//step 4 create a Blockchain
type Blockchain struct {
	Blocks[] *Block // BLocks chain with order
}
 
//step 5 create genesis Block 
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block",[]byte{})
}

//step 6  create a Blockchain with only one genesis block
func newBlockChain() * Blockchain {
	return &Blockchain{ [] *Block{NewGenesisBlock()} }
}

//step 7 add a Block into the Blockchain
func (this *Blockchain) AddBlock(data string){
	//get last Block of Blockchain
	prevBlock :=this.Blocks[len(this.Blocks)-1]
	// create a Block, its prevBlock is last Block of the Blockchain
	newBlock :=NewBlock(data,prevBlock.Hash)
	this.Blocks=append(this.Blocks,newBlock)
}

//step 8 main program , execute the Blockchain process
func main(){
	bc :=newBlockChain()
	var cmd string

	for{
		fmt.Println("Enter 1 to add a data to Blockchain ")
		fmt.Println("Enter 2 to iterate the Blockchain ")
		fmt.Println("Enter otherkey to quit ")
		fmt.Scanf("%s\n", &cmd)

		switch cmd {

			case "1" :
				input :=make([]byte,1024)
				fmt.Println("Please enter Blockchain behavoir data")
				os.Stdin.Read(input)
				bc.AddBlock(string(input))
			case "2":
				for _,block:=range bc.Blocks{
					fmt.Println("===========================")
					fmt.Printf("Prev.Hash: %x\n",block.PrevBlockHash)
					fmt.Printf("Data: %s\n",block.Data)
					fmt.Printf("Hash: %x\n",block.Hash)
					fmt.Println("===========================")
				}
			default:
				fmt.Println("Quit ")
				return
		}
	}

}