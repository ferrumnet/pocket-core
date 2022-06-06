package types

// bridgepool module event types
const (
	EventTransferBySignature    = "transfer_by_signature"
	EventBridgeSwap             = "bridge_swap"
	EventBridgeLiquidityAdded   = "bridge_liquidity_added"
	EventBridgeLiquidityRemoved = "bridge_liquidity_removed"

	AttributeKeySigner        = "signer"
	AttributeKeyReceiver      = "receiver"
	AttributeKeyToken         = "token"
	AttributeKeyAmount        = "amount"
	AttributeKeyFee           = "fee"
	AttributeKeyFrom          = "from"
	AttributeKeyTargetNetwork = "target_network"
	AttributeKeyTargetToken   = "target_token"
	AttributeKeyTargetAddress = "target_address"
	AttributeKeyActor         = "actor"
)
