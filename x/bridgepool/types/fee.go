package types

var (
	TypeMsgSetFee          = "set_fee"
	TypeMsgAllowTarget     = "allow_target"
	TypeMsgDisallowTarget  = "disallow_target"
	TypeMsgAddLiquidity    = "add_liquidity"
	TypeMsgRemoveLiquidity = "remove_liquidity"
	TypeMsgSwap            = "swap"
	TypeMsgWithdrawSigned  = "withdraw_signed"
	TypeMsgAddSigner       = "add_signer"
	TypeMsgRemoveSigner    = "remove_signer"
)

var (
	AppFeeMap = map[string]int64{
		TypeMsgSetFee:          10000,
		TypeMsgAllowTarget:     10000,
		TypeMsgDisallowTarget:  10000,
		TypeMsgAddLiquidity:    10000,
		TypeMsgRemoveLiquidity: 10000,
		TypeMsgSwap:            10000,
		TypeMsgWithdrawSigned:  10000,
		TypeMsgAddSigner:       10000,
		TypeMsgRemoveSigner:    10000,
	}
)
