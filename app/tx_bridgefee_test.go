// nolint
package app

import (
	"testing"
	"time"

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto"
	"github.com/pokt-network/pocket-core/crypto/keys"
	sdk "github.com/pokt-network/pocket-core/types"
	bridgefee "github.com/pokt-network/pocket-core/x/bridgefee"
	bridgefeeTypes "github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/node"
	tmTypes "github.com/tendermint/tendermint/types"
)

func TestBridgeFeeSetTokenInfo(t *testing.T) {
	tt := []struct {
		name         string
		memoryNodeFn func(t *testing.T, genesisState []byte) (tendermint *node.Node, keybase keys.Keybase, cleanup func())
		*upgrades
	}{
		{name: "set with proto codec", memoryNodeFn: NewInMemoryTendermintNodeProto, upgrades: &upgrades{codecUpgrade: codecUpgrade{true, 2}}}, // TODO FULL PROTO SCENARIO
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.upgrades != nil { // NOTE: Use to perform neccesary upgrades for test
				codec.UpgradeHeight = tc.upgrades.codecUpgrade.height
				_ = memCodecMod(tc.upgrades.codecUpgrade.upgradeMod)
			}

			_, kb, cleanup := tc.memoryNodeFn(t, oneAppTwoNodeGenesis())
			time.Sleep(1 * time.Second)
			kp, err := kb.GetCoinbase()
			assert.Nil(t, err)
			_, _, evtChan := subscribeTo(t, tmTypes.EventNewBlock)
			var tx *sdk.TxResponse
			var token = "upokt"
			var bufferSize = uint64(10000)

			<-evtChan // Wait for block
			memCli, stopCli, evtChan := subscribeTo(t, tmTypes.EventTx)
			tx, err = bridgefee.SetTokenInfoTx(memCodec(), memCli, kb, token, bufferSize, kp, "test", tc.codecUpgrade.upgradeMod)
			assert.Nil(t, err)
			assert.NotNil(t, tx)

			select {
			case _ = <-evtChan:
				tokenInfos, err := PCA.QueryBridgeFeeAllTokenInfos(PCA.LastBlockHeight())
				assert.Nil(t, err)
				assert.Equal(t, 0, len(tokenInfos))

				stopCli()
				cleanup()
			}
		})
	}
}

func TestBridgeFeeSetTokenTargetInfos(t *testing.T) {
	tt := []struct {
		name         string
		memoryNodeFn func(t *testing.T, genesisState []byte) (tendermint *node.Node, keybase keys.Keybase, cleanup func())
		*upgrades
	}{
		{name: "set with proto codec", memoryNodeFn: NewInMemoryTendermintNodeProto, upgrades: &upgrades{codecUpgrade: codecUpgrade{true, 2}}}, // TODO FULL PROTO SCENARIO
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.upgrades != nil { // NOTE: Use to perform neccesary upgrades for test
				codec.UpgradeHeight = tc.upgrades.codecUpgrade.height
				_ = memCodecMod(tc.upgrades.codecUpgrade.upgradeMod)
			}

			_, kb, cleanup := tc.memoryNodeFn(t, oneAppTwoNodeGenesis())
			time.Sleep(1 * time.Second)
			kp, err := kb.GetCoinbase()
			assert.Nil(t, err)
			_, _, evtChan := subscribeTo(t, tmTypes.EventNewBlock)
			var tx *sdk.TxResponse
			var token = "upokt"
			var targets = []bridgefeeTypes.TargetInfo{
				{
					Target: sdk.Address(crypto.GenerateEd25519PrivKey().PublicKey().Address()).String(),
					TType:  bridgefeeTypes.TargetType_Address,
					Weight: 1,
				},
			}

			<-evtChan // Wait for block
			memCli, stopCli, evtChan := subscribeTo(t, tmTypes.EventTx)
			tx, err = bridgefee.SetTokenTargetInfosTx(memCodec(), memCli, kb, token, targets, kp, "test", tc.codecUpgrade.upgradeMod)
			assert.Nil(t, err)
			assert.NotNil(t, tx)

			select {
			case _ = <-evtChan:
				targetInfos, err := PCA.QueryBridgeFeeAllTokenTargetInfos(PCA.LastBlockHeight())
				assert.Nil(t, err)
				assert.Equal(t, 0, len(targetInfos))

				stopCli()
				cleanup()
			}
		})
	}
}

func TestBridgeFeeSetGlobalTargetInfos(t *testing.T) {
	tt := []struct {
		name         string
		memoryNodeFn func(t *testing.T, genesisState []byte) (tendermint *node.Node, keybase keys.Keybase, cleanup func())
		*upgrades
	}{
		{name: "set with proto codec", memoryNodeFn: NewInMemoryTendermintNodeProto, upgrades: &upgrades{codecUpgrade: codecUpgrade{true, 2}}}, // TODO FULL PROTO SCENARIO
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.upgrades != nil { // NOTE: Use to perform neccesary upgrades for test
				codec.UpgradeHeight = tc.upgrades.codecUpgrade.height
				_ = memCodecMod(tc.upgrades.codecUpgrade.upgradeMod)
			}

			_, kb, cleanup := tc.memoryNodeFn(t, oneAppTwoNodeGenesis())
			time.Sleep(1 * time.Second)
			kp, err := kb.GetCoinbase()
			assert.Nil(t, err)
			_, _, evtChan := subscribeTo(t, tmTypes.EventNewBlock)
			var tx *sdk.TxResponse
			var targets = []bridgefeeTypes.TargetInfo{
				{
					Target: sdk.Address(crypto.GenerateEd25519PrivKey().PublicKey().Address()).String(),
					TType:  bridgefeeTypes.TargetType_Address,
					Weight: 1,
				},
			}

			<-evtChan // Wait for block
			memCli, stopCli, evtChan := subscribeTo(t, tmTypes.EventTx)
			tx, err = bridgefee.SetGlobalTargetInfosTx(memCodec(), memCli, kb, targets, kp, "test", tc.codecUpgrade.upgradeMod)
			assert.Nil(t, err)
			assert.NotNil(t, tx)
			select {
			case _ = <-evtChan:
				targetInfos, err := PCA.QueryBridgeFeeAllTokenTargetInfos(PCA.LastBlockHeight())
				assert.Nil(t, err)
				assert.Equal(t, 0, len(targetInfos))

				stopCli()
				cleanup()
			}
		})
	}
}
