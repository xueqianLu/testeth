package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"time"
)

func SubScribe(cli *HpbClient, subAddr string) {
	// subscribe newhead
	{
		addr := common.HexToAddress(subAddr)

		newHead := make(chan *types.Header, 20)
		logCh := make(chan types.Log, 20)
		query := ethereum.FilterQuery{
			Addresses: []common.Address{addr},
			Topics:    [][]common.Hash{},
		}
		logsSub, err := cli.SubscribeFilterLogs(context.Background(), query, logCh)
		if err != nil {
			log.Fatal(err)
		}

		sub, err := cli.SubscribeNewHead(context.Background(), newHead)
		if err != nil {
			log.Fatal(err)
		}

		tm := time.NewTicker(time.Second * 3)
		defer tm.Stop()
		for {
			select {
			case <-tm.C:
				blk, err := cli.BlockNumber(context.Background())
				if err != nil {
					log.Fatal(err)
				} else {
					log.Println("get block number ", blk)
				}
			case suberr := <-logsSub.Err():
				log.Fatal(suberr)
			case l, ok := <-logCh:
				if !ok {
					return
				}
				fmt.Println("get new log :", l)
			case suberr := <-sub.Err():
				log.Fatal(suberr)
			case n, ok := <-newHead:
				if !ok {
					return
				}
				fmt.Println("get new header ", n.Number)
			}
		}
	}

}
