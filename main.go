package main

import (
	cmd "demochain/cmd"
	rpc "demochain/rpc"
)

func main(){
	// 以command方式运行
	cmd.Do()
	// 以http方式运行
	rpc.Do()
}

// 用GO语言构建自己的区块链
// https://www.imooc.com/learn/1011