/*
 * @Author: huyongchao huyongchao98@163.com
 * @Date: 2023-02-15 16:14:24
 * @LastEditors: huyongchao huyongchao98@163.com
 * @LastEditTime: 2023-03-30 20:39:35
 * @FilePath: /NearMPCWallet/near-api-go/transaction/transaction_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package transaction

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mr-tron/base58"
	"github.com/near/borsh-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/textileio/near-api-go/keys"
	"github.com/textileio/near-api-go/types"
)

func TestIt(t *testing.T) {
	trans := Transaction{}
	signer, err := keys.NewKeyPairFromRandom("ed25519")
	require.NoError(t, err)
	hash, signedT, err := SignTransaction(trans, signer, "", "")
	require.NoError(t, err)
	require.NotEmpty(t, hash)
	require.NotNil(t, signedT)
	payload, err := borsh.Serialize(*signedT)
	require.NoError(t, err)
	s := base64.StdEncoding.EncodeToString(payload)
	require.NotEmpty(t, s)
}

func TestCreateAccountTransaction(t *testing.T) {
	rpcEndpoint := "https://rpc.testnet.near.org"
	createAccountMethod := "broadcast_tx_async"
	accountID := "adfadfafadfadfafqerqfdafdafdafadfafsdafadgfagagsadfgafadfarewrqerrqreq"
	networkID := "1337"
	_, priv, GenerateKeyErr := ed25519.GenerateKey(nil)

	assert.NoError(t, GenerateKeyErr)

	b58 := base58.Encode(priv)
	keyPair, keyPairErr := keys.NewKeyPairFromString("ed25519:" + b58)

	assert.NoError(t, keyPairErr)

	// 构建JSON-RPC请求
	transaction := NewTransaction(
		accountID,
		keyPair.GetPublicKey(),
		1,
		accountID,
		nil,
		[]Action{CreateAccountAction()},
	)

	hash, st, theErr := SignTransaction(*transaction, keyPair, accountID, networkID)
	assert.NotNil(t, hash)
	assert.NoError(t, theErr)
	jsonRpcRequest := types.JsonRpcRequest{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  createAccountMethod,
		Params:  []interface{}{st},
	}
	requestBody, err := json.Marshal(jsonRpcRequest)
	if err != nil {
		panic(err)
	}

	// 发送JSON-RPC请求
	response, err := http.Post(rpcEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// 解析JSON-RPC响应
	jsonRpcResponse := types.JsonRpcResponse{}
	err = json.NewDecoder(response.Body).Decode(&jsonRpcResponse)
	if err != nil {
		panic(err)
	}

	// 处理JSON-RPC响应
	if jsonRpcResponse.Result != nil {
		result := jsonRpcResponse.Result.(map[string]interface{})
		fmt.Printf("result: %s\n", result["result"])
	} else {
		fmt.Printf("JSON-RPC error: %v\n", jsonRpcResponse)
	}
}
