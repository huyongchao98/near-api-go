/*
 * @Author: huyongchao huyongchao98@163.com
 * @Date: 2023-03-30 10:27:59
 * @LastEditors: huyongchao huyongchao98@163.com
 * @LastEditTime: 2023-04-05 17:43:57
 * @FilePath: /NearMPCWallet/near-api-go/types/jsonRpc.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package types

type JsonRpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type JsonRpcOnlyOneParamsRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Method  string `json:"method"`
	Params  string `json:"params"`
}

type JsonrpcError struct {
	Name    string `json:"name"`
	Cause   Cause  `json:"cause"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Cause struct {
	Name string `json:"name"`
	Info Info   `json:"info"`
}

type Info struct {
	ErrorMessage string `json:"error_message"`
}

type JsonRpcResponse struct {
	Jsonrpc string       `json:"jsonrpc"`
	Id      string       `json:"id"`
	Result  interface{}  `json:"result"`
	Error   JsonrpcError `json:"error"`
}
