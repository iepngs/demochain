package demochain

import core "demochain/core"

func Do(){
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 2 EOS to Jack")
	bc.Print()
}