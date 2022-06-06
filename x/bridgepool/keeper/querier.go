package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for bridgepool REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Ctx, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		default:
			return nil, sdk.ErrUnknownRequest("unknown bridgepool query endpoint")
		}
	}
}
