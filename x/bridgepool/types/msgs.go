package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

var _ sdk.Msg = MsgSetFee{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgSetFee) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgSetFee) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSetFee) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgSetFee) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgSetFee) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgSetFee) Type() string { return TypeMsgSetFee }

// GetFee get fee for msg
func (msg MsgSetFee) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgAllowTarget{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAllowTarget) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgAllowTarget) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAllowTarget) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgAllowTarget) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgAllowTarget) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgAllowTarget) Type() string { return TypeMsgAllowTarget }

// GetFee get fee for msg
func (msg MsgAllowTarget) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgDisallowTarget{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgDisallowTarget) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgDisallowTarget) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgDisallowTarget) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgDisallowTarget) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgDisallowTarget) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgDisallowTarget) Type() string { return TypeMsgDisallowTarget }

// GetFee get fee for msg
func (msg MsgDisallowTarget) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgAddLiquidity{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddLiquidity) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgAddLiquidity) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAddLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgAddLiquidity) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgAddLiquidity) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgAddLiquidity) Type() string { return TypeMsgAddLiquidity }

// GetFee get fee for msg
func (msg MsgAddLiquidity) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgRemoveLiquidity{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgRemoveLiquidity) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgRemoveLiquidity) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgRemoveLiquidity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgRemoveLiquidity) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgRemoveLiquidity) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgRemoveLiquidity) Type() string { return TypeMsgRemoveLiquidity }

// GetFee get fee for msg
func (msg MsgRemoveLiquidity) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgSwap{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgSwap) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgSwap) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSwap) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgSwap) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgSwap) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgSwap) Type() string { return TypeMsgSwap }

// GetFee get fee for msg
func (msg MsgSwap) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgWithdrawSigned{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgWithdrawSigned) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgWithdrawSigned) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgWithdrawSigned) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgWithdrawSigned) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgWithdrawSigned) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgWithdrawSigned) Type() string { return TypeMsgWithdrawSigned }

// GetFee get fee for msg
func (msg MsgWithdrawSigned) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgAddSigner{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddSigner) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgAddSigner) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAddSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgAddSigner) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgAddSigner) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgAddSigner) Type() string { return TypeMsgAddSigner }

// GetFee get fee for msg
func (msg MsgAddSigner) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgRemoveSigner{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgRemoveSigner) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgRemoveSigner) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgRemoveSigner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgRemoveSigner) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgRemoveSigner) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgRemoveSigner) Type() string { return TypeMsgRemoveSigner }

// GetFee get fee for msg
func (msg MsgRemoveSigner) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}
