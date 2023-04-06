/*
 * @Author: huyongchao huyongchao98@163.com
 * @Date: 2023-03-16 15:29:00
 * @LastEditors: huyongchao huyongchao98@163.com
 * @LastEditTime: 2023-04-06 20:37:01
 * @FilePath: /NearMPCWallet/near-api-go/account/test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/textileio/near-api-go/types"
)

type WalletCreateParams struct {
	SeedPhrase string `json:"seedPhrase"`
	PublicKey  string `json:"publicKey"`
}

func TestAccount(t *testing.T) {
	rpcEndpoint := "https://rpc.testnet.near.org"
	walletCreateMethod := "near_wallet_create"

	// 构建JSON-RPC请求
	walletCreateParams := WalletCreateParams{}
	jsonRpcRequest := types.JsonRpcRequest{
		Jsonrpc: "2.0",
		Id:      "1",
		Method:  walletCreateMethod,
		Params:  []interface{}{walletCreateParams},
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
		fmt.Printf("Seed phrase: %s\n", result["seed_phrase"])
		fmt.Printf("Public key: %s\n", result["public_key"])
		fmt.Printf("Account ID: %s\n", result["account_id"])
	} else {
		fmt.Printf("JSON-RPC error: %v\n", jsonRpcResponse)
	}
}
