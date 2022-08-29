package types

// bridgepool module event types
const (
	EventTransferBySignature    = "transfer_by_signature"
	EventBridgeSwap             = "bridge_swap"
	EventBridgeLiquidityAdded   = "bridge_liquidity_added"
	EventBridgeLiquidityRemoved = "bridge_liquidity_removed"
	EventTypeSetFeeRate         = "set_fee_rate"
	EventTypeAllowTarget        = "allow_target"
	EventTypeDisallowTarget     = "disallow_target"
	EventTypeSetSigner          = "set_signer"
	EventTypeRemoveSigner       = "remove_signer"

	AttributeKeySigner        = "signer"
	AttributeKeyReceiver      = "receiver"
	AttributeKeyToken         = "token"
	AttributeKeyAmount        = "amount"
	AttributeKeyFee           = "fee"
	AttributeKeyFrom          = "from"
	AttributeKeyTargetChainId = "target_chain_id"
	AttributeKeyTargetToken   = "target_token"
	AttributeKeyTargetAddress = "target_address"
	AttributeKeyActor         = "actor"
	AttributeKeyChainId       = "chain_id"
)
