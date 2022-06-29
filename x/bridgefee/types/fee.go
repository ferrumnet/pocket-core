package types

var (
	TypeMsgAddAllowedActor      = "add_allowed_actor"
	TypeMsgRemoveAllowedActor   = "remove_allowed_actor"
	TypeMsgSetTokenInfo         = "set_token_info"
	TypeMsgSetTokenTargetInfos  = "set_token_target_infos"
	TypeMsgSetGlobalTargetInfos = "set_global_target_infos"
)

var (
	AppFeeMap = map[string]int64{
		TypeMsgAddAllowedActor:      10000,
		TypeMsgRemoveAllowedActor:   10000,
		TypeMsgSetTokenInfo:         10000,
		TypeMsgSetTokenTargetInfos:  10000,
		TypeMsgSetGlobalTargetInfos: 10000,
	}
)
