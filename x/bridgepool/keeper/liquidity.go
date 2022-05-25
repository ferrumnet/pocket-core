package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) SetLiquidity(ctx sdk.Ctx, token string, user string, amount uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.Cdc.MarshalBinaryBare(amount, ctx.BlockHeight())
	if err != nil {
		return err
	}
	return store.Set(types.LiquidityKey(token, user), bz)
}

func (k Keeper) AddLiquidity(ctx sdk.Ctx, token string, user string, amount uint64) error {
	// TODO: send tokens from the user to module `amount`

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeLiquidityAdded,
			sdk.NewAttribute(types.AttributeKeyActor, user),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	liquidity := k.GetLiquidity(ctx, token, user)
	return k.SetLiquidity(ctx, token, user, liquidity+amount)
}

func (k Keeper) RemoveLiquidity(ctx sdk.Ctx, token string, user string, amount uint64) error {
	// TODO: send tokens from the module to user for `amount`

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeLiquidityRemoved,
			sdk.NewAttribute(types.AttributeKeyActor, user),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	liquidity := k.GetLiquidity(ctx, token, user)
	if liquidity < amount {
		return types.NotEnoughLiquidityError
	}

	return k.SetLiquidity(ctx, token, user, liquidity-amount)
}

func (k Keeper) GetLiquidity(ctx sdk.Ctx, token string, user string) uint64 {
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
