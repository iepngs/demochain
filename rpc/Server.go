package demochain

import (
	core "demochain/core"
	"encoding/json"
	"io"
	"net/http"
)

var blockchain *core.Blockchain

func run() {
	http.HandleFunc("/blockchain/get", blockchainGetHandler)
	http.HandleFunc("/blockchain/write", blockchainWriteHandler)
	http.ListenAndServe(":8085", nil)
}

func blockchainGetHandler(w http.ResponseWriter, r * http.Request){
	bytes, err := json.Marshal(blockchain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter, r * http.Request){
	blockData := r.URL.Query().Get("data")
	blockchain.SendData(blockData)
	blockchainGetHandler(w, r)
}

func Do(){
	blockchain = core.NewBlockchain()
	run()
}