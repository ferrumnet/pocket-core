package types

const (
	ModuleName   = "bridgefee" // name of the module
	RouterKey    = ModuleName  // RouterKey defines the routing key for a Parameter Change
	StoreKey     = ModuleName  // key for state store
	TStoreKey    = "transient_bridgefee"
	QuerierRoute = ModuleName // QuerierRoute is the querier route for the bridgefee module
)

var (
	TokenInfoKeyPrefix       = []byte{0x01}
	TokenTargetInfoKeyPrefix = []byte{0x02}
)

func TokenInfoKey(token string) []byte {
	return append(TokenInfoKeyPrefix, token...)
}

func TokenTargetInfoKey(token string) []byte {
	return append(TokenTargetInfoKeyPrefix, token...)
}
