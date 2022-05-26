package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) SetLiquidity(ctx sdk.Ctx, token string, user sdk.Address, amount uint64) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.Cdc.MarshalBinaryBare(amount, ctx.BlockHeight())
	if err != nil {
		panic(err)
	}
	store.Set(types.LiquidityKey(token, user), bz)
	return nil
}

func (k Keeper) AddLiquidity(ctx sdk.Ctx, token string, user sdk.Address, amount uint64) sdk.Error {
	// send tokens from the user to module
	err := k.AccountKeeper.SendCoinsFromAccountToModule(ctx, user, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(token, int64(amount))})
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeLiquidityAdded,
			sdk.NewAttribute(types.AttributeKeyActor, user.String()),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	liquidity := k.GetLiquidity(ctx, token, user)
	return k.SetLiquidity(ctx, token, user, liquidity+amount)
}

func (k Keeper) RemoveLiquidity(ctx sdk.Ctx, token string, user sdk.Address, amount uint64) sdk.Error {
	// send tokens from the module to user for `amount`
	err := k.AccountKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user, sdk.Coins{sdk.NewInt64Coin(token, int64(amount))})
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeLiquidityRemoved,
			sdk.NewAttribute(types.AttributeKeyActor, user.String()),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	liquidity := k.GetLiquidity(ctx, token, user)
	if liquidity < amount {
		return types.ErrNotEnoughLiquidity(k.codespace)
	}

	return k.SetLiquidity(ctx, token, user, liquidity-amount)
}

func (k Keeper) GetLiquidity(ctx sdk.Ctx, token string, user sdk.Address) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.LiquidityKey(token, user))
	if err != nil {
		return 0
	}
	if bz == nil {
		return 0
	}
	amount := uint64(0)
	err = k.Cdc.UnmarshalBinaryBare(bz, &amount, ctx.BlockHeight())
	if err != nil {
		return 0
	}
	return amount
}
