package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"sync"
	"time"
)

func StressRpc(cli *HpbClient, n int) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for m := 0; m < 100000; m++ {
				num, err := cli.BlockNumber(ctx)
				if err != nil {
					log.Info("get block ", "routines ", i, "err", err)
				} else {
					log.Info("get block ", "routines ", i, "block", num)
				}
				time.Sleep(time.Millisecond * 200)
			}
		}(i)
	}
	wg.Wait()
}
