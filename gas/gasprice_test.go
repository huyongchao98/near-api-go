package gas

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	"github.com/textileio/near-api-go/types"
	"testing"
)

var ctx = context.Background()

func TestGasPrice_ViewState(t *testing.T) {
	rpcClient, err := rpc.DialContext(ctx, "https://rpc.testnet.near.org")
	require.NoError(t, err)

	config := &types.Config{
		RPCClient: rpcClient,
		NetworkID: "testnet",
	}
	type fields struct {
		config *types.Config
	}
	type args struct {
		ctx         context.Context
		blockHeight interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GasPriceStateView
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:   "TestGasPrice_ViewState",
			fields: fields{config: config},
			args:   args{ctx: ctx, blockHeight: "58pz1fwQEbyMU8nCts4UmzrDH3t9ECMuzmtUxpL8gvcc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewGasPrice(tt.fields.config)
			got, err := a.ViewState(tt.args.ctx, tt.args.blockHeight)
			if nil != err {
				t.Errorf("ViewState() error = %v", err)
			} else {
				t.Logf("ViewState() got = %v, want %v", got, tt.want)
			}
		})
	}
}
