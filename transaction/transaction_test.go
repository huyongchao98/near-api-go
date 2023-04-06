/*
 * @Author: huyongchao huyongchao98@163.com
 * @Date: 2023-02-15 16:14:24
 * @LastEditors: huyongchao huyongchao98@163.com
 * @LastEditTime: 2023-04-06 20:28:23
 * @FilePath: /NearMPCWallet/near-api-go/transaction/transaction_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package transaction

import (
	"encoding/base64"
	"testing"

	"github.com/near/borsh-go"
	"github.com/stretchr/testify/require"
	"github.com/textileio/near-api-go/keys"
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
