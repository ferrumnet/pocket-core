package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace       sdk.CodespaceType = ModuleName
	CodeUnexpectedError    CodeType          = 101
	CodeNotEnoughLiquidity CodeType          = 102
)

func ErrNotEnoughLiquidity(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeNotEnoughLiquidity, "not enough liquidity")
}

func ErrUnexpectedError(Codespace sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(Codespace, CodeUnexpectedError, err.Error())
}
