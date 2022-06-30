package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
)

func (k Keeper) AddAllowedActor(ctx sdk.Ctx, actor string) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.AllowedActorKey(actor), []byte(actor))
	return nil
}

func (k Keeper) RemoveAllowedActor(ctx sdk.Ctx, actor string) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AllowedActorKey(actor))
	return nil
}

func (k Keeper) IsAllowedActor(ctx sdk.Ctx, actor string) bool {
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.AllowedActorKey(actor))
	if bz == nil || err != nil {
		return false
	}
	return true
}

func (k Keeper) GetAllAllowedActors(ctx sdk.Ctx) []string {
	actors := []string{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.AllowedActorKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		actors = append(actors, string(iterator.Value()))
	}
	return actors
}

func (k Keeper) SetTokenInfo(ctx sdk.Ctx, info types.TokenInfo) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	bz := k.Cdc.MustMarshalJSON(&info)
	store.Set(types.TokenInfoKey(info.Token), bz)
	return nil
}

func (k Keeper) GetTokenInfo(ctx sdk.Ctx, token string) types.TokenInfo {
	info := types.TokenInfo{}
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.TokenInfoKey(token))
	if err != nil || bz == nil {
		return info
	}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info
}

func (k Keeper) GetAllTokenInfos(ctx sdk.Ctx, token string) []types.TokenInfo {
	infos := []types.TokenInfo{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.TokenInfoKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		info := types.TokenInfo{}
		k.Cdc.MustUnmarshalJSON(iterator.Value(), &info)
		infos = append(infos, info)
	}
	return infos
}

func (k Keeper) SetTokenTargetInfo(ctx sdk.Ctx, info types.TokenTargetInfo) sdk.Error {
	store := ctx.KVStore(k.storeKey)
	bz := k.Cdc.MustMarshalJSON(&info)
	store.Set(types.TokenTargetInfoKey(info.Token), bz)
	return nil
}

func (k Keeper) GetTokenTargetInfo(ctx sdk.Ctx, token string) types.TokenTargetInfo {
	info := types.TokenTargetInfo{}
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.TokenTargetInfoKey(token))
	if err != nil || bz == nil {
		return info
	}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info
}

func (k Keeper) GetGlobalTokenTargetInfo(ctx sdk.Context) types.TokenTargetInfo {
	info := types.TokenTargetInfo{}
	store := ctx.KVStore(k.storeKey)
	bz, err := store.Get(types.TokenTargetInfoKey(""))
	if err != nil || bz == nil {
		return info
	}
	k.Cdc.MustUnmarshalJSON(bz, &info)
	return info
}

func (k Keeper) GetAllTokenTargetInfos(ctx sdk.Ctx, token string) []types.TokenTargetInfo {
	infos := []types.TokenTargetInfo{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.TokenTargetInfoKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		info := types.TokenTargetInfo{}
		k.Cdc.MustUnmarshalJSON(iterator.Value(), &info)
		infos = append(infos, info)
	}
	return infos
}
