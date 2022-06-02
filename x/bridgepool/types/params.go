package types

import (
	"reflect"

	sdk "github.com/pokt-network/pocket-core/types"
)

// Parameter keys
var (
	DefaultParamspace = ModuleName

	OwnerKey = []byte("owner")
)

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
// pairs of bridgepool module's parameters.
// nolint
func (p *Params) ParamSetPairs() sdk.ParamSetPairs {
	return sdk.ParamSetPairs{
		{Key: OwnerKey, Value: &p.Owner},
	}
}

// Equal returns a boolean determining if two Params types are identical.
func (p Params) Equal(p2 Params) bool {
	return reflect.DeepEqual(p, p2)
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		Owner: "",
	}
}
