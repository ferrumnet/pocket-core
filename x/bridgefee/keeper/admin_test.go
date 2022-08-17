package keeper

import (
	"testing"

	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/stretchr/testify/assert"
)

func TestTokenInfo(t *testing.T) {
	context, _, keeper := createTestInput(t, true)

	// check initial token infos
	tokenInfo := keeper.GetTokenInfo(context, "upokt")
	assert.Equal(t, tokenInfo, types.TokenInfo{})

	tokenInfos := keeper.GetAllTokenInfos(context)
	assert.Len(t, tokenInfos, 0)

	// set token info
	info := types.TokenInfo{
		Token:               "upokt",
		BufferSize:          1000000,
		TokenSpecificConfig: 1,
	}
	err := keeper.SetTokenInfo(context, info)
	assert.Nil(t, err)

	// check token info
	tokenInfo = keeper.GetTokenInfo(context, "upokt")
	assert.Equal(t, tokenInfo, info)

	// check all token infos
	tokenInfos = keeper.GetAllTokenInfos(context)
	assert.Len(t, tokenInfos, 1)

	// set token info for second token
	info2 := types.TokenInfo{
		Token:               "upokt2",
		BufferSize:          1000000,
		TokenSpecificConfig: 1,
	}
	err = keeper.SetTokenInfo(context, info2)
	assert.Nil(t, err)

	// check all token infos
	tokenInfos = keeper.GetAllTokenInfos(context)
	assert.Len(t, tokenInfos, 2)
}

func TestTokenTargetInfo(t *testing.T) {
	context, _, keeper := createTestInput(t, true)

	// check initial token infos
	targetInfo := keeper.GetTokenTargetInfo(context, "upokt")
	assert.Equal(t, targetInfo, types.TokenTargetInfo{})

	targetInfo = keeper.GetGlobalTokenTargetInfo(context)
	assert.Equal(t, targetInfo, types.TokenTargetInfo{})

	targetInfos := keeper.GetAllTokenTargetInfos(context)
	assert.Len(t, targetInfos, 0)

	privKey1 := crypto.GenerateEd25519PrivKey()
	pubKey1 := privKey1.PublicKey()
	addr1 := sdk.Address(pubKey1.Address())
	privKey2 := crypto.GenerateEd25519PrivKey()
	pubKey2 := privKey2.PublicKey()
	addr2 := sdk.Address(pubKey2.Address())
	// set token info
	info := types.TokenTargetInfo{
		Token: "upokt",
		Targets: []types.TargetInfo{
			{
				Target: addr1.String(),
				TType:  types.TargetType_Burn,
				Weight: 1,
			},
			{
				Target: addr2.String(),
				TType:  types.TargetType_Address,
				Weight: 1,
			},
		},
	}
	err := keeper.SetTokenTargetInfo(context, info)
	assert.Nil(t, err)

	// check token info
	targetInfo = keeper.GetTokenTargetInfo(context, "upokt")
	assert.Equal(t, targetInfo, info)

	// check all token infos
	targetInfos = keeper.GetAllTokenTargetInfos(context)
	assert.Len(t, targetInfos, 1)

	// set token info for global target
	info2 := types.TokenTargetInfo{
		Token: "",
		Targets: []types.TargetInfo{
			{
				Target: addr1.String(),
				TType:  types.TargetType_Burn,
				Weight: 1,
			},
			{
				Target: addr2.String(),
				TType:  types.TargetType_Address,
				Weight: 1,
			},
		},
	}
	err = keeper.SetTokenTargetInfo(context, info2)
	assert.Nil(t, err)

	targetInfo = keeper.GetGlobalTokenTargetInfo(context)
	assert.Equal(t, targetInfo, info2)

	// check all token infos
	targetInfos = keeper.GetAllTokenTargetInfos(context)
	assert.Len(t, targetInfos, 2)
}
