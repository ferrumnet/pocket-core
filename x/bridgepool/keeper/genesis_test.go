package keeper

import (
	"testing"

	"github.com/pokt-network/pocket-core/x/bridgepool/types"
	"github.com/stretchr/testify/assert"
)

func TestExportGenesis(t *testing.T) {
	context, _, keeper := createTestInput(t, true)

	defaultGenesis := types.DefaultGenesisState()
	genesis := keeper.ExportGenesis(context)
	assert.Equal(t, genesis, defaultGenesis)
}

func TestInitGenesis(t *testing.T) {
	context, _, keeper := createTestInput(t, true)

	genesis := types.DefaultGenesisState()
	genesis.AllowedTargets = []types.AllowedTarget{{
		"upokt", "137", "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc",
	}}
	genesis.Signers = []string{
		"0x88848510E92B2f2F7ea06d46e7B370198F7369Bc",
	}
	genesis.Liquidities = []types.Liquidity{
		{
			"upokt", "98848510e92b2f2f7ea06d46e7b370198f7369bc", 10000,
		},
	}
	genesis.Fees = []types.FeeRate{
		{
			"upokt", 100,
		},
	}
	genesis.UsedWithdrawMessages = [][]byte{[]byte("example")}
	keeper.InitGenesis(context, genesis)

	signers := keeper.GetAllSigners(context)
	assert.Equal(t, signers, genesis.Signers)
	liquidities := keeper.GetAllLiquidities(context)
	assert.Equal(t, liquidities, genesis.Liquidities)
	feeRates := keeper.GetAllFeeRates(context)
	assert.Equal(t, feeRates, genesis.Fees)
	allowedTargets := keeper.GetAllAllowedTargets(context)
	assert.Equal(t, allowedTargets, genesis.AllowedTargets)
	usedMessages := keeper.GetAllUsedMessages(context)
	assert.Equal(t, usedMessages, genesis.UsedWithdrawMessages)
}
