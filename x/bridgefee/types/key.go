package types

const (
	ModuleName = "bridgefee" // name of the module
	RouterKey  = ModuleName  // RouterKey defines the routing key for a Parameter Change
	StoreKey   = ModuleName  // key for state store
)

var (
	AllowedActorKeyPrefix    = []byte{0x00}
	TokenInfoKeyPrefix       = []byte{0x01}
	TokenTargetInfoKeyPrefix = []byte{0x02}
)

func AllowedActorKey(actor string) []byte {
	return append(AllowedActorKeyPrefix, actor...)
}

func TokenInfoKey(token string) []byte {
	return append(TokenInfoKeyPrefix, token...)
}

func TokenTargetInfoKey(token string) []byte {
	return append(TokenTargetInfoKeyPrefix, token...)
}
