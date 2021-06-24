package main

import (
	"context"
	"flag"
	"time"
)

func main() {
	url := flag.String("u", "http://127.0.0.1:8545", "rpc url")
	subAddr := flag.String("addr", "", "subscribe address")
	routines := flag.Int("n", 0, "routines number that stress test used ")
	lrandom := flag.Bool("random", false, "loop get random")
	tduration := flag.Int("d", 10, "test duration")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(*tduration))
	defer cancel()

	cli, err := CreateHpbClient(*url)
	if err != nil {
		panic(err)
	}

	if *routines > 0 {
		StressRpc(cli, *routines)
	}

	if len(*subAddr) > 0 {
		SubScribe(cli, *subAddr)
	}

	if *lrandom {
		RandomTest(ctx, cli)
	}
}
