package types

// bridgefee module event types
const (
	EventTypeAddAllowedActor           = "add_allowed_actor"
	EventTypeRemoveAllowedActor        = "remove_allowed_actor"
	EventTypeSetTokenTargetInfos       = "set_token_target_infos"
	EventTypeSetGlobalTokenTargetInfos = "set_global_token_target_infos"

	AttributeKeyActor  = "actor"
	AttributeKeySender = "sender"
	AttributeKeyToken  = "token"
)
