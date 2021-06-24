package main

import (
	"context"
	"log"
	"math/big"
	"time"
)

func RandomTest(ctx context.Context, cli *HpbClient) {
	var curBlock uint64
	tc := time.NewTicker(time.Second * 2)
	defer tc.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-tc.C:
			newBlock, err := cli.BlockNumber(ctx)
			if err != nil {
				log.Println("backend get block number failed, err ", err)
				return
			}
			if newBlock > curBlock {
				curBlock = newBlock
				rdm, err := cli.GetRandom(ctx, big.NewInt(int64(curBlock)))
				if err != nil {
					log.Println("get random failed, err ", err)
					return
				}
				log.Println("Get random ", rdm)
			}
		}
	}
}
