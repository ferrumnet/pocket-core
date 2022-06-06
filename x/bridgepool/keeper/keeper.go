package keeper

import (
	"github.com/pokt-network/pocket-core/codec"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	AccountKeeper types.AccountKeeper
	Paramstore    sdk.Subspace
	storeKey      sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc           *codec.Codec // The wire codec for binary encoding/decoding.
	// codespace
	codespace sdk.CodespaceType
}

// NewKeeper creates new instances of the pocketcore module Keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, accountKeeper types.AccountKeeper, paramstore sdk.Subspace, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		AccountKeeper: accountKeeper,
		Paramstore:    paramstore.WithKeyTable(ParamKeyTable()),
		storeKey:      storeKey,
		Cdc:           cdc,
		codespace:     codespace,
	}
}

func (k Keeper) Codec() *codec.Codec {
	return k.Cdc
}
