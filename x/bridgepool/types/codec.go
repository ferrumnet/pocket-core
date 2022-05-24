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
	cdc.RegisterStructure(MsgSetFee{}, "bridgepool/msg_set_fee")
	cdc.RegisterStructure(MsgAllowTarget{}, "bridgepool/msg_allow_target")
	cdc.RegisterStructure(MsgDisallowTarget{}, "bridgepool/msg_disallow_target")
	cdc.RegisterStructure(MsgAddLiquidity{}, "bridgepool/msg_add_liquidity")
	cdc.RegisterStructure(MsgRemoveLiquidity{}, "bridgepool/msg_remove_liquidity")
	cdc.RegisterStructure(MsgSwap{}, "bridgepool/msg_swap")
	cdc.RegisterStructure(MsgWithdrawSigned{}, "bridgepool/msg_withdraw_signed")

	cdc.RegisterInterface("x.interface.nil", (*interface{})(nil))

	cdc.RegisterImplementation((*sdk.ProtoMsg)(nil),
		&MsgSetFee{},
		&MsgAllowTarget{},
		&MsgDisallowTarget{},
		&MsgAddLiquidity{},
		&MsgRemoveLiquidity{},
		&MsgSwap{},
		&MsgWithdrawSigned{},
	)
	cdc.RegisterImplementation((*sdk.Msg)(nil),
		&MsgSetFee{},
		&MsgAllowTarget{},
		&MsgDisallowTarget{},
		&MsgAddLiquidity{},
		&MsgRemoveLiquidity{},
		&MsgSwap{},
		&MsgWithdrawSigned{},
	)
	ModuleCdc = cdc
}
