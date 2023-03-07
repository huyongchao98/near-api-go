package gas

import (
	"context"
	"fmt"
	"github.com/textileio/near-api-go/types"
	"github.com/textileio/near-api-go/util"
)

type GasPrice struct {
	config *types.Config
}

type GasPriceStateView struct {
	GasPrice string `json:"gas_price"`
}

func NewGasPrice(config *types.Config) *GasPrice {
	return &GasPrice{
		config: config,
	}
}

// ViewState queries the gas_price.
// @args block_height(...int) or block_hash(...string)
func (a *GasPrice) ViewState(ctx context.Context, args ...interface{}) (*GasPriceStateView, error) {
	var res GasPriceStateView
	if err := a.config.RPCClient.CallContext(ctx, &res, "gas_price", args...); err != nil {
		return nil, fmt.Errorf("calling rpc: %v", util.MapRPCError(err))
	}
	return &res, nil
}
