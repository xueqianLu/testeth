package main

import (
	"flag"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	url := flag.String("u", "http://127.0.0.1:8545", "rpc url")
	subAddr := flag.String("addr", "", "subscribe address")
	routines := flag.Int("n", 0, "routines number that stress test used ")
	flag.Parse()

	cli, err := ethclient.Dial(*url)
	if err != nil {
		panic(err)
	}

	if *routines > 0 {
		StressRpc(cli, *routines)
	}

	if len(*subAddr) > 0 {
		SubScribe(cli, *subAddr)
	}
}
