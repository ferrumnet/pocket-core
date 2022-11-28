package keeper

import (
	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

// SetFeeRate set fee rate for a token on the store
func (k Keeper) SetFeeRate(ctx sdk.Ctx, token string, fee10000 uint64) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	info := types.FeeRate{
		Token: token,
		Rate:  fee10000,
	}
	bz := k.Cdc.MustMarshalJSON(&info)
	store.Set(types.FeeRateKey(token), bz)
	return nil
}

// GetFeeRate gets fee rate for a token from the store
func (k Keeper) GetFeeRate(ctx sdk.Ctx, token string) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.FeeRateKey(token))
	if err != nil {
		return 0
	}
	if bz == nil {
		return 0
	}
	info := types.FeeRate{}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info.Rate
}

// GetFeeRate gets all registered fee rates from the store
func (k Keeper) GetAllFeeRates(ctx sdk.Ctx) []types.FeeRate {
	feeRates := []types.FeeRate{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.FeeRateKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		info := types.FeeRate{}
		k.Cdc.MustUnmarshalJSON(iterator.Value(), &info)
		feeRates = append(feeRates, info)
	}
	return feeRates
}

// AllowTarget allow swap target chain and target token address for a denom
func (k Keeper) AllowTarget(ctx sdk.Ctx, token string, chainId string, targetToken string) sdk.Error {
	// check ethereum addresses
	if !common.IsHexAddress(targetToken) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}

	store := ctx.KVStore(k.storeKey)
	info := types.AllowedTarget{
		Token:       token,
		ChainId:     chainId,
		TargetToken: targetToken,
	}
	bz := k.Cdc.MustMarshalJSON(&info)
	store.Set(types.AllowedTargetKey(token, chainId), bz)
	return nil
}

// DisallowTarget disallow swap target chain for a denom
func (k Keeper) DisallowTarget(ctx sdk.Ctx, token string, chainId string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AllowedTargetKey(token, chainId))
}

// GetAllowedTarget gets target chain token address for a denom
func (k Keeper) GetAllowedTarget(ctx sdk.Ctx, token string, chainId string) string {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.AllowedTargetKey(token, chainId))
	if err != nil || bz == nil {
		return ""
	}
	info := types.AllowedTarget{}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info.TargetToken
}

// GetAllAllowedTargets gets all allowed targets
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

// SetSigner set allowed signer by module admin
func (k Keeper) SetSigner(ctx sdk.Ctx, signer string) sdk.Error {
	// check ethereum addresses
	if !common.IsHexAddress(signer) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.SignerKey(signer), []byte(signer))
	return nil
}

// DeleteSigner removes allowed signer by module admin
func (k Keeper) DeleteSigner(ctx sdk.Ctx, signer string) sdk.Error {
	// check ethereum addresses
	if !common.IsHexAddress(signer) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SignerKey(signer))
	return nil
}

// IsSigner returns if an address is signer or not
func (k Keeper) IsSigner(ctx sdk.Ctx, signer string) bool {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.SignerKey(signer))
	if bz == nil || err != nil {
		return false
	}
	return true
}

// GetAllSigners returns all registered signers
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
