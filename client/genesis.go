package client

import (
	"context"

	"github.com/eteu-technologies/near-api-go/jsonrpc"
)

// TODO: decode response
// https://docs.near.org/docs/develop/front-end/rpc#genesis-config
func (c *Client) GenesisConfig(ctx context.Context) (res jsonrpc.JSONRPCResponse, err error) {
	res, err = c.doRPC(ctx, "EXPERIMENTAL_genesis_config", nil, nil)

	return
}
