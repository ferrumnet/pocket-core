package keeper

import (
	"testing"

	"github.com/pokt-network/pocket-core/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	apptypes "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/stretchr/testify/assert"
)

func genAccount() sdk.Address {
	privKey := crypto.GenerateEd25519PrivKey()
	pubKey := privKey.PublicKey()
	return sdk.Address(pubKey.Address())
}

func TestLiquidityRawStorage(t *testing.T) {
	addr1 := genAccount()
	addr2 := genAccount()

	context, _, keeper, _ := createTestInput(t, true)

	// check liquidity when not set
	liquidity := keeper.GetLiquidity(context, "upokt", addr1)
	assert.Equal(t, liquidity, uint64(0))

	liquidities := keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 0)

	// set liquidity
	err := keeper.SetLiquidity(context, "upokt", addr1, 100)
	assert.Nil(t, err)

	// check if liquidity is set correctly
	liquidity = keeper.GetLiquidity(context, "upokt", addr1)
	assert.Equal(t, liquidity, uint64(100))

	// check all liquidities
	liquidities = keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 1)

	// set second liquidity
	err = keeper.SetLiquidity(context, "upokt", addr2, 200)
	assert.Nil(t, err)

	// set second liquidity
	err = keeper.SetLiquidity(context, "upokt2", addr1, 100)
	assert.Nil(t, err)

	// check all liquidities
	liquidities = keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 3)
}

func TestLiquidityAddRemove(t *testing.T) {
	addr1 := genAccount()
	addr2 := genAccount()

	context, _, keeper, _ := createTestInput(t, true)

	// check liquidity when not set
	liquidity := keeper.GetLiquidity(context, "upokt", addr1)
	assert.Equal(t, liquidity, uint64(0))

	liquidities := keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 0)

	// add liquidity when account does not have enough balance
	err := keeper.AddLiquidity(context, "upokt", addr1, 100)
	assert.NotNil(t, err)

	// prepare coins on addr1
	coins := sdk.Coins{sdk.NewInt64Coin("upokt", 1000), sdk.NewInt64Coin("upokt2", 1000)}
	err = keeper.AccountKeeper.MintCoins(context, apptypes.StakedPoolName, coins)
	assert.Nil(t, err)
	err = keeper.AccountKeeper.MintCoins(context, apptypes.StakedPoolName, coins)
	assert.Nil(t, err)
	err = keeper.AccountKeeper.SendCoinsFromModuleToAccount(context, apptypes.StakedPoolName, addr1, coins)
	assert.Nil(t, err)
	err = keeper.AccountKeeper.SendCoinsFromModuleToAccount(context, apptypes.StakedPoolName, addr2, coins)
	assert.Nil(t, err)

	// add liquidity when account has enough balance
	err = keeper.AddLiquidity(context, "upokt", addr1, 100)
	assert.Nil(t, err)

	// check if liquidity is set correctly
	liquidity = keeper.GetLiquidity(context, "upokt", addr1)
	assert.Equal(t, liquidity, uint64(100))

	// check all liquidities
	liquidities = keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 1)

	// add second account liquidity
	err = keeper.AddLiquidity(context, "upokt2", addr2, 200)
	assert.Nil(t, err)

	// remove liquidity liquidity
	err = keeper.RemoveLiquidity(context, "upokt", addr1, 100)
	assert.Nil(t, err)

	// check all liquidities
	liquidities = keeper.GetAllLiquidities(context)
	assert.Len(t, liquidities, 2)
}
