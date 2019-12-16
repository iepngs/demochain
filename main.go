package main

import (
	cmd "demochain/cmd"
	rpc "demochain/rpc"
)

func main(){
	cmd.Do()
	rpc.Do()
}

// 用GO语言构建自己的区块链
// https://www.imooc.com/learn/1011