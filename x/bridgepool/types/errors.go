package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace               sdk.CodespaceType = ModuleName
	CodeUnexpectedError            CodeType          = 101
	CodeNotEnoughLiquidity         CodeType          = 102
	CodeInvalidSignature           CodeType          = 103
	CodeInvalidSigner              CodeType          = 104
	CodeNotEnoughPermission        CodeType          = 105
	CodeAlreadyUsedWithdrawMessage CodeType          = 106
	CodeInvalidEthereumAddress     CodeType          = 107
)

func ErrNotEnoughPermission(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeNotEnoughPermission, "not enough permission")
}

func ErrNotEnoughLiquidity(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeNotEnoughLiquidity, "not enough liquidity")
}

func ErrInvalidSignature(Codespace sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidSignature, err.Error())
}

func ErrInvalidSigner(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidSigner, "Invalid signer")
}

func ErrInvalidEthereumAddress(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeInvalidEthereumAddress, "Invalid ethereum address")
}

func ErrAlreadyUsedWithdrawMessage(Codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(Codespace, CodeAlreadyUsedWithdrawMessage, "Already used withdraw message")
}

func ErrUnexpectedError(Codespace sdk.CodespaceType, err error) sdk.Error {
	return sdk.NewError(Codespace, CodeUnexpectedError, err.Error())
}
