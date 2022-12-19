package main

//func init() {
//	log.SetPrefix("Blockchain: ")
//}

func main() {
	bc := NewBlockchain()

	bc.AddTransaction("A", "B", 1.0)
	bc.AddTransaction("x", "z", 5.0)
	bc.AddTransaction("x", "A", 7.1)
	lastBlockHash := bc.LastBlock().Hash256()
	bc.AddBlock(5, lastBlockHash)

	bc.AddTransaction("c", "g", 5.0)
	bc.AddTransaction("c", "f", 7.1)
	lastBlockHash = bc.LastBlock().Hash256()
	bc.AddBlock(7, lastBlockHash)

	bc.AddTransaction("e", "b", 5.0)
	lastBlockHash = bc.LastBlock().Hash256()
	bc.AddBlock(7, lastBlockHash)

	bc.Print()
}
