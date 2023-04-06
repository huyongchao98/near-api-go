package account

import (
	"context"
	"fmt"
	"math/big"

	"testing"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	"github.com/textileio/near-api-go/keys"
	"github.com/textileio/near-api-go/transaction"
	"github.com/textileio/near-api-go/types"
)

var ctx = context.Background()

func TestIt(t *testing.T) {
	a, cleanup := makeAccount(t)
	defer cleanup()
	require.NotNil(t, a)
}

// func TestViewState(t *testing.T) {
// 	a, cleanup := makeAccount(t)
// 	defer cleanup()
// 	res, err := a.ViewState(ctx, ViewStateWithFinality("final"))
// 	require.NoError(t, err)
// 	require.NotNil(t, res)
// }

// func TestState(t *testing.T) {
// 	a, cleanup := makeAccount(t)
// 	defer cleanup()
// 	res, err := a.State(ctx, StateWithFinality("final"))
// 	require.NoError(t, err)
// 	require.NotNil(t, res)
// }

// func TestFindAccessKey(t *testing.T) {
// 	a, cleanup := makeAccount(t)
// 	defer cleanup()
// 	pubKey, accessKeyView, err := a.FindAccessKey(ctx, "", nil)
// 	require.NoError(t, err)
// 	require.NotNil(t, pubKey)
// 	require.NotNil(t, accessKeyView)
// }

var PublicKey = "ed25519:H9k5eiU4xXS3M4z8HzKJSLaZdqGdGwBG49o7orNC4eZW"
var accountID = "client.chainlink.testnet"

// func TestCreateAccount(t *testing.T) {
// 	rpcClient, err := rpc.DialContext(ctx, "https://rpc.testnet.near.org")
// 	require.NoError(t, err)

// 	config := &types.Config{
// 		RPCClient: rpcClient,
// 		NetworkID: "testnet",
// 		// Signer:    keys,
// 	}
// 	a := NewAccount(config, accountID)
// 	creatAccountErr := a.CreateAccount(PublicKey)
// 	require.NoError(t, creatAccountErr)

// }

func TestViewAccessKey(t *testing.T) {
	a, cleanup := makeAccount(t)
	defer cleanup()

	publicKey := &keys.PublicKey{
		Type: 0,
		Data: []byte(PublicKey),
	}
	v, err := a.ViewAccessKey(ctx, publicKey)
	require.NoError(t, err)
	require.NotNil(t, v)
}

func TestSignTransaction(t *testing.T) {
	a, cleanup := makeAccount(t)
	defer cleanup()
	amt := big.NewInt(1000)
	sendAction := transaction.TransferAction(*amt)
	hash, signedTxn, err := a.SignTransaction(ctx, "carsonfarmer.testnet", sendAction)
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.NotNil(t, signedTxn)
}

func TestSignAndSendTransaction(t *testing.T) {
	a, cleanup := makeAccount(t)
	defer cleanup()
	amt, ok := (&big.Int{}).SetString("1000000000000000000000000", 10)
	require.True(t, ok)
	sendAction := transaction.TransferAction(*amt)
	res, err := a.SignAndSendTransaction(ctx, "carsonfarmer.testnet", sendAction)
	require.NoError(t, err)
	require.NotNil(t, res)

	status, ok := res.GetStatus()
	fmt.Println(status, ok)

	status2, ok := res.GetStatusBasic()
	fmt.Println(status2, ok)
}

func getKeyPair() (keys.KeyPair, error) {
	keypair, err := keys.NewKeyPairFromRandom("ED25519")
	publicKey := keypair.GetPublicKey()
	publicKeyString, toStringErr := publicKey.ToString()
	fmt.Println("publicKey:", publicKeyString)
	fmt.Println("toStringErr:", toStringErr)
	return keypair, err
	// return keys.NewKeyPairFromString(
	// 	"ed25519:H9k5eiU4xXS3M4z8HzKJSLaZdqGdGwBG49o7orNC4eZW",
	// )
}

func makeAccount(t *testing.T) (*Account, func()) {
	rpcClient, err := rpc.DialContext(ctx, "https://rpc.testnet.near.org")
	require.NoError(t, err)

	keys, err := getKeyPair()
	require.NoError(t, err)

	config := &types.Config{
		RPCClient: rpcClient,
		NetworkID: "testnet",
		Signer:    keys,
	}
	a := NewAccount(config, "2423423423adaf.testnet")
	return a, func() {
		rpcClient.Close()
	}
}

func TestCreateAccountTransaction(t *testing.T) {
	rpcEndpoint := "https://rpc.testnet.near.org"
	mainAccountID := "testdafa.testnet"
	accountID := "example-account3.testnet"
	networkID := "testnet"
	// _, priv, GenerateKeyErr := ed25519.GenerateKey(nil)

	// assert.NoError(t, GenerateKeyErr)

	// b58 := base58.Encode(priv)

	keyPair, keyPairErr := keys.NewKeyPairFromString("ed25519:3G7BmuSTuo825Y1kCTyRwMm9incjuNDcf24p42pKi9PgDv3JyvPzJT4Kb88mRHR3KyPDXNu2Gsy3w8dRMAR6eKoM")
	require.NoError(t, keyPairErr)

	publicKey := keyPair.GetPublicKey()
	publicKeyString, toStringErr := publicKey.ToString()
	fmt.Println("publicKey:", publicKeyString)
	fmt.Println("toStringErr:", toStringErr)

	rpcClient, rpcClientErr := rpc.DialContext(ctx, rpcEndpoint)
	require.NoError(t, rpcClientErr)
	config := &types.Config{
		Signer:    keyPair,
		NetworkID: networkID,
		RPCClient: rpcClient,
	}
	theMainAccount := NewAccount(config, mainAccountID)

	//发起创建账号请求
	finalExecutionOutcome, transactionErr := theMainAccount.SignAndSendTransaction(ctx, accountID, transaction.CreateAccountAction())

	require.NoError(t, transactionErr, "报错了")

	_, success := finalExecutionOutcome.GetStatus()

	require.True(t, success)

}
