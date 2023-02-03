package main

import (
	"go-blockchain/blockchain"
	"go-blockchain/wallet"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*blockchain.Blockchain = make(map[string]*blockchain.Blockchain)

type BlockchainServer struct {
	port uint16
}

func NewServer(p uint16) *BlockchainServer {
	return &BlockchainServer{p}
}

func (s *BlockchainServer) Port() uint16 {
	return s.port
}

func (s *BlockchainServer) GetBlockchain() *blockchain.Blockchain {
	bc, exist := cache["blockchain"]
	if !exist {
		minerWallet := wallet.NewWallet()
		bc = blockchain.NewBlockchain(minerWallet.GetAddress(), s.Port())
		cache["blockchain"] = bc

		log.Printf("Private key: %v", minerWallet.GetPrivateKey())
		log.Printf("Public key: %v", minerWallet.GetPublicKey())
		log.Printf("Address: %v", minerWallet.GetAddress())
	}

	return bc
}

func (s *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := s.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("Err: wrong methode")
	}
}

func (s *BlockchainServer) Run() {
	http.HandleFunc("/", s.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(s.port)), nil))
}

func Ping(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Pong!")
}
