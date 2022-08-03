package types

import (
	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/codec/types"
	sdk "github.com/pokt-network/pocket-core/types"
)

// module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.NewCodec(types.NewInterfaceRegistry())
	RegisterCodec(ModuleCdc)
	ModuleCdc.AminoCodec().Seal()
}

// RegisterCodec registers all necessary param module types with a given codec.
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterStructure(MsgSetTokenInfo{}, "bridgefee/msg_set_token_info")
	cdc.RegisterStructure(MsgSetTokenTargetInfos{}, "bridgefee/msg_set_token_target_infos")
	cdc.RegisterStructure(MsgSetGlobalTargetInfos{}, "bridgefee/msg_set_global_target_infos")

	cdc.RegisterImplementation((*sdk.ProtoMsg)(nil),
		&MsgSetTokenInfo{},
		&MsgSetTokenTargetInfos{},
		&MsgSetGlobalTargetInfos{},
	)
	cdc.RegisterImplementation((*sdk.Msg)(nil),
		&MsgSetTokenInfo{},
		&MsgSetTokenTargetInfos{},
		&MsgSetGlobalTargetInfos{},
	)
	ModuleCdc = cdc
}
