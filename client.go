package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

const (
	MaxIdleConns        int = 100
	MaxIdleConnsPerHost int = 100
	IdleConnTimeout     int = 40
)

type HpbClient struct {
	c *rpc.Client
	*ethclient.Client
	url string
}

func CreateHpbClient(url string) (*HpbClient, error) {
	c, err := rpc.DialContext(context.Background(), url)
	if err != nil {
		return nil, err
	}
	ethclient, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	return &HpbClient{c, ethclient, url}, nil
}
func toBlockNumArg(number *big.Int) string {
	if number == nil {
		return "latest"
	}
	pending := big.NewInt(-1)
	if number.Cmp(pending) == 0 {
		return "pending"
	}
	return hexutil.EncodeBig(number)
}
func (hc *HpbClient) GetRandom(ctx context.Context, blockNumber *big.Int) (string, error) {
	var raw json.RawMessage
	err := hc.c.CallContext(ctx, &raw, "hpb_getRandom", toBlockNumArg(blockNumber))
	if err != nil {
		return "", err
	} else if len(raw) == 0 {
		return "", ethereum.NotFound
	}
	fmt.Println("got random hex:", hex.EncodeToString(raw), "string:", string(raw))
	return string(raw), nil
}
