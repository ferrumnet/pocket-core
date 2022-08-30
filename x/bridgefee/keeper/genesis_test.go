package keeper

import (
	"testing"

	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
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
	genesis.Params = types.Params{
		Owner: sdk.Address(crypto.GenerateEd25519PrivKey().PublicKey().Address()).String(),
	}
	genesis.TokenInfos = []types.TokenInfo{{
		"upokt", 1000, 1,
	}}

	genesis.TokenTargets = []types.TokenTargetInfo{{
		"upokt", []types.TargetInfo{
			{
				Target: sdk.Address(crypto.GenerateEd25519PrivKey().PublicKey().Address()).String(),
				TType:  types.TargetType_Address,
				Weight: 1,
			},
		},
	}}

	keeper.InitGenesis(context, genesis)
	params := keeper.GetParams(context)
	assert.Equal(t, params, genesis.Params)
	tokenInfos := keeper.GetAllTokenInfos(context)
	assert.Equal(t, tokenInfos, genesis.TokenInfos)
	targetInfos := keeper.GetAllTokenTargetInfos(context)
	assert.Equal(t, targetInfos, genesis.TokenTargets)
}
