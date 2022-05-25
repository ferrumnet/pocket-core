package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) SetFeeRate(ctx sdk.Ctx, token string, fee10000 uint64) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.Cdc.MarshalBinaryBare(fee10000, ctx.BlockHeight())
	if err != nil {
		return err
	}
	return store.Set(types.FeeRateKey(token), bz)
}

func (k Keeper) GetFeeRate(ctx sdk.Ctx, token string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.FeeRateKey(token))
	if err != nil {
		return 0
	}
	if bz == nil {
		return 0
	}
	fee10000 := uint64(0)
	err = k.Cdc.UnmarshalBinaryBare(bz, &fee10000, ctx.BlockHeight())
	if err != nil {
		return 0
	}
	return fee10000
}

func (k Keeper) AllowTarget(ctx sdk.Ctx, token string, chainId uint64, targetToken string) error {
	store := ctx.KVStore(k.storeKey)
	return store.Set(types.AllowedTargetKey(token, chainId), []byte(targetToken))
}

func (k Keeper) DisallowTarget(ctx sdk.Ctx, token string, chainId uint64) error {
	store := ctx.KVStore(k.storeKey)
	return store.Delete(types.AllowedTargetKey(token, chainId))
}

func (k Keeper) GetAllowedTarget(ctx sdk.Ctx, token string, chainId uint64) string {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.AllowedTargetKey(token, chainId))
	if err != nil {
		return ""
	}
	return string(bz)
}
