package keeper

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/stretchr/testify/assert"
)

func TestGetSigner(t *testing.T) {
	tests := []struct {
		name      string
		payee     string
		amount    sdk.Coin
		salt      string
		signer    string
		signature string
		errors    bool
	}{
		{
			name:      "get signer test",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			amount:    sdk.NewInt64Coin("upokt", 100000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc",
			errors:    false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			signer, _, err := GetSigner("testnet", tc.payee, tc.amount, tc.salt, common.Hex2Bytes(tc.signature))
			assert.Nil(t, err)
			assert.Equal(t, signer.String(), tc.signer)
		})
	}
}

func TestWithdrawSigned(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		payee     string
		amount    uint64
		salt      string
		signature string
		errors    bool
	}{
		{
			name:      "withdraw pass",
			token:     "0xA719b8aB7EA7AF0DDb4358719a34631bb79d15Dc",
			payee:     "0x65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			amount:    100, //30431598848000000000000,
			salt:      "0x74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "0xfd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			errors:    false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			context, accs, keeper := createTestInput(t, true)
			_, _, _ = context, accs, keeper
			// signer, err := GetSigner(tc.token, tc.payee, tc.amount, common.Hex2Bytes(tc.salt), common.Hex2Bytes(tc.signature))
			// fmt.Println("signer", signer.String())
			// assert.NotNil(t, err)
			// err := keeper.WithdrawSigned(context, accs[0].String(), tc.token, tc.payee, tc.amount, common.Hex2Bytes(tc.salt), common.Hex2Bytes(tc.signature))
			// switch tc.errors {
			// case true:
			// 	assert.NotNil(t, err)
			// default:

			// }
		})
	}
}
