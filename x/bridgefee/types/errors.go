package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace        sdk.CodespaceType = ModuleName
	CodeNotEnoughPermission CodeType          = 101
)

func ErrNotEnoughPermission(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeNotEnoughPermission, "not enough permission")
}
