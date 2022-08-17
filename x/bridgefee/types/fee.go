package types

var (
	TypeMsgSetTokenInfo         = "set_token_info"
	TypeMsgSetTokenTargetInfos  = "set_token_target_infos"
	TypeMsgSetGlobalTargetInfos = "set_global_target_infos"
)

var (
	AppFeeMap = map[string]int64{
		TypeMsgSetTokenInfo:         10000,
		TypeMsgSetTokenTargetInfos:  10000,
		TypeMsgSetGlobalTargetInfos: 10000,
	}
)
