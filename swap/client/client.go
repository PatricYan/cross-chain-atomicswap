package client

import (
	"context"
	"math/big"

	"github.com/seeleteam/go-seele/common"
	rpc "github.com/seeleteam/go-seele/rpc2"
)

// Client defines typed wrappers for the Ethereum ABI RPC API.
type Client struct {
	c *rpc.Client
}

// NewClient create a client
func NewClient(addressValue string) (*Client, error) {
	client, err := rpc.DialTCP(context.Background(), addressValue)
	return &Client{client}, err
}

// CodeAt returns the code of the given account.
func (client *Client) CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error) {
	// var result hexutil.Bytes
	// err := client.c.Call(&result, "seele_getCode",)
	// return result, err
	return nil, nil
}
