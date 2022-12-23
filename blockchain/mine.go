package blockchain

import "log"

const MiningDifficulty int = 2
const MiningSender = "Main Chain"
const MiningReward = 25

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MiningSender, bc.address, MiningReward, nil, nil)
	nonce := bc.GenerateProofOfWorkValidNonce()
	ph := bc.LastBlock().Hash256()
	bc.AddBlock(nonce, ph)
	log.Println("action=mining, status=success")

	return true
}
