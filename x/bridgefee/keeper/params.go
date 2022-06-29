package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
)

// "ParamKeyTable" - Registers the paramset types in a keytable and returns the table
func ParamKeyTable() sdk.KeyTable {
	return sdk.NewKeyTable().RegisterParamSet(&types.Params{})
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Ctx) types.Params {
	return types.Params{}
}

// set the params
func (k Keeper) SetParams(ctx sdk.Ctx, params types.Params) {
	k.Paramstore.SetParamSet(ctx, &params)
}
