package keeper

import (
	"testing"

	sdk "github.com/pokt-network/pocket-core/types"
	apptypes "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
	"github.com/stretchr/testify/assert"
)

func TestSwap(t *testing.T) {
	addr1 := genAccount()

	context, _, keeper, _ := createTestInput(t, true)

	// try swap for not allowed target
	err := keeper.Swap(context, addr1, "upokt", 1000000, "137", "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc", "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.NotNil(t, err)

	// allow target
	err = keeper.AllowTarget(context, "upokt", "137", "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// deposit balance
	coins := sdk.Coins{sdk.NewInt64Coin("upokt", 1000000)}
	err = keeper.AccountKeeper.MintCoins(context, apptypes.StakedPoolName, coins)
	assert.Nil(t, err)
	err = keeper.AccountKeeper.SendCoinsFromModuleToAccount(context, apptypes.StakedPoolName, addr1, coins)
	assert.Nil(t, err)

	oldAddr1Balance := keeper.AccountKeeper.GetCoins(context, addr1)
	oldModuleBalance := keeper.AccountKeeper.GetCoins(context, keeper.AccountKeeper.GetModuleAddress(types.ModuleName))

	err = keeper.Swap(context, addr1, "upokt", 1000000, "137", "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc", "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	newAddr1Balance := keeper.AccountKeeper.GetCoins(context, addr1)
	newModuleBalance := keeper.AccountKeeper.GetCoins(context, keeper.AccountKeeper.GetModuleAddress(types.ModuleName))

	swapCoin := sdk.NewInt64Coin("upokt", 1000000)
	assert.Equal(t, newModuleBalance, oldModuleBalance.Add(sdk.Coins{swapCoin}))
	assert.Equal(t, oldAddr1Balance, newAddr1Balance.Add(sdk.Coins{swapCoin}))
}
