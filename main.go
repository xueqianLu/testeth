package main

import(
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {
	url := flag.String("u", "http://127.0.0.1:8545", "rpc url")
	flag.Parse()
	cli,err := ethclient.Dial(*url)
	if err != nil {
		panic(err)
	}

	newHead := make(chan *types.Header, 20)
	sub, err := cli.SubscribeNewHead(context.Background(), newHead)
	if err != nil {
		log.Fatal(err)
	}

	tm := time.NewTicker(time.Second * 3)
	defer tm.Stop()
	for {
		select {
		case <- tm.C:
			blk, err := cli.BlockNumber(context.Background())
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("get block number ", blk)
			}
		case suberr := <- sub.Err():
			log.Fatal(suberr)
		case n,ok := <- newHead:
			if !ok {
				return
			}
			fmt.Println("get new header ", n.Number)
		}
	}
}

