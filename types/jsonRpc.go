/*
 * @Author: huyongchao huyongchao98@163.com
 * @Date: 2023-03-30 10:27:59
 * @LastEditors: huyongchao huyongchao98@163.com
 * @LastEditTime: 2023-03-30 10:29:03
 * @FilePath: /NearMPCWallet/near-api-go/types/jsonRpc.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package types

type JsonRpcRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Id      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type JsonRpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  interface{} `json:"result"`
}
