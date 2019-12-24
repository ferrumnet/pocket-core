package keeper

import (
	"github.com/pokt-network/pocket-core/x/pocketcore/types"
	"github.com/pokt-network/posmint/codec"
	"github.com/pokt-network/posmint/crypto/keys"
	sdk "github.com/pokt-network/posmint/types"
	"github.com/pokt-network/posmint/x/params"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	posKeeper          types.PosKeeper
	appKeeper          types.AppsKeeper
	Keybase            *keys.Keybase
	coinbasePassphrase string // todo is this safe??
	hostedBlockchains  types.HostedBlockchains
	Paramstore         params.Subspace
	storeKey           sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc                *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewPocketCoreKeeper creates new instances of the pocketcore Keeper
func NewPocketCoreKeeper(storeKey sdk.StoreKey, cdc *codec.Codec, posKeeper types.PosKeeper, appKeeper types.AppsKeeper, hostedChains types.HostedBlockchains, paramstore params.Subspace, passphrase string) Keeper {
	return Keeper{
		storeKey:           storeKey,
		cdc:                cdc,
		posKeeper:          posKeeper,
		appKeeper:          appKeeper,
		coinbasePassphrase: passphrase,
		hostedBlockchains:  hostedChains,
		Paramstore:         paramstore.WithKeyTable(ParamKeyTable()),
	}
}