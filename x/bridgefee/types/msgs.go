package types

import (
	sdk "github.com/pokt-network/pocket-core/types"
)

var _ sdk.Msg = MsgAddAllowedActor{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddAllowedActor) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgAddAllowedActor) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgAddAllowedActor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgAddAllowedActor) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgAddAllowedActor) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgAddAllowedActor) Type() string { return TypeMsgAddAllowedActor }

// GetFee get fee for msg
func (msg MsgAddAllowedActor) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgRemoveAllowedActor{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgRemoveAllowedActor) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgRemoveAllowedActor) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgRemoveAllowedActor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgRemoveAllowedActor) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgRemoveAllowedActor) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgRemoveAllowedActor) Type() string { return TypeMsgRemoveAllowedActor }

// GetFee get fee for msg
func (msg MsgRemoveAllowedActor) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgSetTokenInfo{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgSetTokenInfo) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgSetTokenInfo) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSetTokenInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgSetTokenInfo) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgSetTokenInfo) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgSetTokenInfo) Type() string { return TypeMsgSetTokenInfo }

// GetFee get fee for msg
func (msg MsgSetTokenInfo) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgSetTokenTargetInfos{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgSetTokenTargetInfos) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgSetTokenTargetInfos) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSetTokenTargetInfos) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgSetTokenTargetInfos) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgSetTokenTargetInfos) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgSetTokenTargetInfos) Type() string { return TypeMsgSetTokenTargetInfos }

// GetFee get fee for msg
func (msg MsgSetTokenTargetInfos) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}

var _ sdk.Msg = MsgSetGlobalTargetInfos{}

// GetSigners return address(es) that must sign over msg.GetSignBytes()
func (msg MsgSetGlobalTargetInfos) GetSigners() []sdk.Address {
	return []sdk.Address{msg.FromAddress}
}

func (msg MsgSetGlobalTargetInfos) GetRecipient() sdk.Address {
	return nil
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSetGlobalTargetInfos) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic quick validity check for staking an application
func (msg MsgSetGlobalTargetInfos) ValidateBasic() sdk.Error {
	return nil
}

// Route provides router key for msg
func (msg MsgSetGlobalTargetInfos) Route() string { return RouterKey }

// Type provides msg name
func (msg MsgSetGlobalTargetInfos) Type() string { return TypeMsgSetGlobalTargetInfos }

// GetFee get fee for msg
func (msg MsgSetGlobalTargetInfos) GetFee() sdk.BigInt {
	return sdk.NewInt(AppFeeMap[msg.Type()])
}
