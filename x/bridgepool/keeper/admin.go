package keeper

import (
	"bytes"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) SetFeeRate(ctx sdk.Ctx, token string, fee10000 uint64) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.Cdc.MarshalBinaryBare(fee10000, ctx.BlockHeight())
	if err != nil {
		panic(err)
	}
	store.Set(types.FeeRateKey(token), bz)
	return nil
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

func (k Keeper) GetAllFeeRates(ctx sdk.Ctx) []types.FeeRate {
	feeRates := []types.FeeRate{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.FeeRateKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		tokenBytes := bytes.TrimPrefix(iterator.Key(), types.FeeRateKeyPrefix)
		fee10000 := uint64(0)
		err := k.Cdc.UnmarshalBinaryBare(iterator.Value(), &fee10000, ctx.BlockHeight())
		if err != nil {
			panic(err)
		}
		feeRates = append(feeRates, types.FeeRate{
			Token: string(tokenBytes),
			Rate:  fee10000,
		})
	}
	return feeRates
}

func (k Keeper) AllowTarget(ctx sdk.Ctx, token string, chainId uint64, targetToken string) {
	store := ctx.KVStore(k.storeKey)
	info := types.AllowedTarget{
		Token:       token,
		ChainId:     chainId,
		TargetToken: targetToken,
	}
	bz := k.Cdc.MustMarshalJSON(&info)
	store.Set(types.AllowedTargetKey(token, chainId), bz)
}

func (k Keeper) DisallowTarget(ctx sdk.Ctx, token string, chainId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AllowedTargetKey(token, chainId))
}

func (k Keeper) GetAllowedTarget(ctx sdk.Ctx, token string, chainId uint64) string {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.AllowedTargetKey(token, chainId))
	if err != nil {
		return ""
	}
	info := types.AllowedTarget{}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info.TargetToken
}

func (k Keeper) GetAllAllowedTargets(ctx sdk.Ctx) []types.AllowedTarget {
	allowedTargets := []types.AllowedTarget{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.AllowedTargetKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		info := types.AllowedTarget{}
		k.Cdc.MustUnmarshalJSON(iterator.Value(), &info)
		allowedTargets = append(allowedTargets, info)
	}
	return allowedTargets
}

func (k Keeper) SetSigner(ctx sdk.Ctx, signer string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SignerKey(signer), []byte(signer))
}

func (k Keeper) DeleteSigner(ctx sdk.Ctx, signer string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SignerKey(signer))
}

func (k Keeper) IsSigner(ctx sdk.Ctx, signer string) bool {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.SignerKey(signer))
	if bz == nil || err != nil {
		return false
	}
	return true
}

func (k Keeper) GetAllSigners(ctx sdk.Ctx) []string {
	signers := []string{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.SignerKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		signers = append(signers, string(iterator.Value()))
	}
	return signers
}
