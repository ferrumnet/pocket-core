package keeper

import (
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

func (k Keeper) GetAllFeeRates(ctx sdk.Ctx) {
	// TODO: implement

	// 	applications = make([]types.Application, 0)
	// store := ctx.KVStore(k.storeKey)
	// iterator, _ := sdk.KVStorePrefixIterator(store, types.AllApplicationsKey)
	// defer iterator.Close()

	// for ; iterator.Valid(); iterator.Next() {
	// 	application, err := types.UnmarshalApplication(k.Cdc, ctx, iterator.Value())
	// 	if err != nil {
	// 		k.Logger(ctx).Error("couldn't unmarshal application in GetAllApplications call: " + string(iterator.Value()) + "\n" + err.Error())
	// 		continue
	// 	}
	// 	applications = append(applications, application)
	// }
	// return applications
}

func (k Keeper) AllowTarget(ctx sdk.Ctx, token string, chainId uint64, targetToken string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedTargetKey(token, chainId), []byte(targetToken))
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
	return string(bz)
}

func (k Keeper) GetAllAllowedTargets(ctx sdk.Ctx) {
	// TODO: implement

	// 	applications = make([]types.Application, 0)
	// store := ctx.KVStore(k.storeKey)
	// iterator, _ := sdk.KVStorePrefixIterator(store, types.AllApplicationsKey)
	// defer iterator.Close()

	// for ; iterator.Valid(); iterator.Next() {
	// 	application, err := types.UnmarshalApplication(k.Cdc, ctx, iterator.Value())
	// 	if err != nil {
	// 		k.Logger(ctx).Error("couldn't unmarshal application in GetAllApplications call: " + string(iterator.Value()) + "\n" + err.Error())
	// 		continue
	// 	}
	// 	applications = append(applications, application)
	// }
	// return applications
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
	// TODO: implement

	// 	applications = make([]types.Application, 0)
	// store := ctx.KVStore(k.storeKey)
	// iterator, _ := sdk.KVStorePrefixIterator(store, types.AllApplicationsKey)
	// defer iterator.Close()

	// for ; iterator.Valid(); iterator.Next() {
	// 	application, err := types.UnmarshalApplication(k.Cdc, ctx, iterator.Value())
	// 	if err != nil {
	// 		k.Logger(ctx).Error("couldn't unmarshal application in GetAllApplications call: " + string(iterator.Value()) + "\n" + err.Error())
	// 		continue
	// 	}
	// 	applications = append(applications, application)
	// }
	// return applications
	return []string{}
}
