package keeper

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	apptypes "github.com/pokt-network/pocket-core/x/apps/types"
	bridgefeeTypes "github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
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
	privKey1 := crypto.GenerateEd25519PrivKey()
	pubKey1 := privKey1.PublicKey()
	addr1 := sdk.Address(pubKey1.Address())

	tests := []struct {
		name      string
		payee     string
		feeRate   uint64
		moduleBal sdk.Coin
		amount    sdk.Coin
		salt      string
		signature string
		signer    string
		feeTarget sdk.Address
		errors    bool
	}{
		{
			name:      "withdraw pass",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			feeRate:   0,
			moduleBal: sdk.NewInt64Coin("upokt", 100000),
			amount:    sdk.NewInt64Coin("upokt", 100000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc",
			feeTarget: addr1,
			errors:    false,
		},
		{
			name:      "Fee non-zero case",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			feeRate:   1000, // 10%
			moduleBal: sdk.NewInt64Coin("upokt", 100000),
			amount:    sdk.NewInt64Coin("upokt", 100000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x77777777E92B2f2F7ea06d46e7B370198F7369Bc",
			feeTarget: sdk.Address{},
			errors:    false,
		},
		{
			name:      "Fee distribution set case",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			feeRate:   1000, // 10%
			moduleBal: sdk.NewInt64Coin("upokt", 100000),
			amount:    sdk.NewInt64Coin("upokt", 100000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x77777777E92B2f2F7ea06d46e7B370198F7369Bc",
			feeTarget: addr1,
			errors:    false,
		},
		{
			name:      "invalid signer case",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			feeRate:   0,
			moduleBal: sdk.NewInt64Coin("upokt", 100000),
			amount:    sdk.NewInt64Coin("upokt", 100000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x77777777E92B2f2F7ea06d46e7B370198F7369Bc",
			feeTarget: addr1,
			errors:    false, // TODO: should be true when signer checker is set
		},
		{
			name:      "Not enough balance on module case",
			payee:     "65A8F07Bd9A8598E1b5B6C0a88F4779DBC077675",
			feeRate:   0,
			moduleBal: sdk.NewInt64Coin("upokt", 100000),
			amount:    sdk.NewInt64Coin("upokt", 10000000000000),
			salt:      "74b82d4c386d401a708f401b5f1b831cb7c34adc2f02f39e60a4d5d220d66303",
			signature: "fd41b43ee89d28652b34f0e5c769753ba3b4f0cb6b85bc63f29ba6680cedbf89688d66145b5326c1afb811117a196f0f14b73989b062dde980ce13cd06f1ad801b",
			signer:    "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc",
			feeTarget: addr1,
			errors:    true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			context, accs, keeper, feeKeeper := createTestInput(t, true)
			_, _, _ = context, accs, keeper
			err := keeper.SetFeeRate(context, tc.amount.Denom, tc.feeRate)
			assert.Nil(t, err)
			err = keeper.AccountKeeper.MintCoins(context, apptypes.StakedPoolName, sdk.Coins{tc.moduleBal})
			assert.Nil(t, err)
			err = keeper.AccountKeeper.SendCoinsFromModuleToModule(context, apptypes.StakedPoolName, types.ModuleName, sdk.Coins{tc.moduleBal})
			assert.Nil(t, err)

			if tc.feeTarget.String() != "" {
				feeKeeper.SetTokenInfo(context, bridgefeeTypes.TokenInfo{
					Token:               "upokt",
					BufferSize:          1,
					TokenSpecificConfig: 1,
				})
				feeKeeper.SetTokenTargetInfo(context, bridgefeeTypes.TokenTargetInfo{
					Token: "upokt",
					Targets: []bridgefeeTypes.TargetInfo{
						{
							Target: addr1.String(),
							TType:  bridgefeeTypes.TargetType_Address,
							Weight: 1,
						},
					},
				})
			}

			feePoolAddr := keeper.AccountKeeper.GetModuleAddress(bridgefeeTypes.ModuleName)
			oldFeepoolBal := keeper.AccountKeeper.GetCoins(context, feePoolAddr)
			err = keeper.WithdrawSigned(context, accs[0].String(), tc.payee, tc.amount, tc.salt, common.Hex2Bytes(tc.signature))

			switch tc.errors {
			case true:
				assert.NotNil(t, err)
			default:
				assert.Nil(t, err)
				if tc.feeRate > 0 {
					if tc.feeTarget.String() != "" {
						feeTargetBal := keeper.AccountKeeper.GetCoins(context, tc.feeTarget)
						assert.True(t, feeTargetBal.AmountOf(tc.amount.Denom).IsPositive())
					} else {
						newFeepoolBal := keeper.AccountKeeper.GetCoins(context, feePoolAddr)
						assert.True(t, newFeepoolBal.AmountOf(tc.amount.Denom).GT(oldFeepoolBal.AmountOf(tc.amount.Denom)))
					}
				}
				// try withdraw signed again
				err = keeper.WithdrawSigned(context, accs[0].String(), tc.payee, tc.amount, tc.salt, common.Hex2Bytes(tc.signature))
				assert.NotNil(t, err)
			}
		})
	}
}
