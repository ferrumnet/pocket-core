package keeper

import (
	"testing"

	types2 "github.com/pokt-network/pocket-core/codec/types"

	"github.com/pokt-network/pocket-core/crypto"
	"github.com/pokt-network/pocket-core/types/module"
	"github.com/pokt-network/pocket-core/x/gov"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/store"
	sdk "github.com/pokt-network/pocket-core/types"
	apptypes "github.com/pokt-network/pocket-core/x/apps/types"
	"github.com/pokt-network/pocket-core/x/auth"
	bridgefeeKeeper "github.com/pokt-network/pocket-core/x/bridgefee/keeper"
	bridgefeeTypes "github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

var (
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		gov.AppModuleBasic{},
	)
)

// : deadcode unused
// create a codec used only for testing
func makeTestCodec() *codec.Codec {
	var cdc = codec.NewCodec(types2.NewInterfaceRegistry())
	auth.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	crypto.RegisterAmino(cdc.AminoCodec().Amino)
	return cdc
}

// : deadcode unused
func createTestInput(t *testing.T, isCheckTx bool) (sdk.Context, []auth.Account, Keeper, bridgefeeKeeper.Keeper) {
	initPower := int64(100000000000)
	nAccs := int64(4)

	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.ParamsKey
	tkeyParams := sdk.ParamsTKey
	keyBridgepool := sdk.NewKVStoreKey(types.ModuleName)
	keyBridgefee := sdk.NewKVStoreKey(bridgefeeTypes.ModuleName)
	db := dbm.NewMemDB()
	ms := store.NewCommitMultiStore(db, false, 5000000)
	ms.MountStoreWithDB(keyAcc, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBridgepool, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyBridgefee, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(keyParams, sdk.StoreTypeIAVL, db)
	ms.MountStoreWithDB(tkeyParams, sdk.StoreTypeTransient, db)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	ctx := sdk.NewContext(ms, abci.Header{ChainID: "test-chain"}, isCheckTx, log.NewNopLogger()).WithAppVersion("0.0.0")
	ctx = ctx.WithConsensusParams(
		&abci.ConsensusParams{
			Validator: &abci.ValidatorParams{
				PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519},
			},
		},
	)
	cdc := makeTestCodec()

	maccPerms := map[string][]string{
		apptypes.StakedPoolName:   {auth.Minter, auth.Burner},
		auth.FeeCollectorName:     nil,
		types.ModuleName:          nil,
		bridgefeeTypes.ModuleName: {auth.Burner},
	}
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[auth.NewModuleAddress(acc).String()] = true
	}
	valTokens := sdk.TokensFromConsensusPower(initPower)
	accSubspace := sdk.NewSubspace(auth.DefaultParamspace)
	bridgepoolSubspace := sdk.NewSubspace(types.ModuleName)
	bridgefeeSubspace := sdk.NewSubspace(bridgefeeTypes.ModuleName)
	ak := auth.NewKeeper(cdc, keyAcc, accSubspace, maccPerms)
	moduleManager := module.NewManager(
		auth.NewAppModule(ak),
	)
	genesisState := ModuleBasics.DefaultGenesis()
	moduleManager.InitGenesis(ctx, genesisState)
	initialCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultStakeDenom, valTokens))
	accs := createTestAccs(ctx, int(nAccs), initialCoins, &ak)
	feeKeeper := bridgefeeKeeper.NewKeeper(keyBridgefee, cdc, ak, bridgefeeSubspace, "bridgefee")
	keeper := NewKeeper(keyBridgepool, cdc, ak, feeKeeper, bridgepoolSubspace, "bridgepool")
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)
	return ctx, accs, keeper, feeKeeper
}

// : unparam deadcode unused
func createTestAccs(ctx sdk.Ctx, numAccs int, initialCoins sdk.Coins, ak *auth.Keeper) (accs []auth.Account) {
	for i := 0; i < numAccs; i++ {
		privKey := crypto.GenerateEd25519PrivKey()
		pubKey := privKey.PublicKey()
		addr := sdk.Address(pubKey.Address())
		acc := auth.NewBaseAccountWithAddress(addr)
		acc.Coins = initialCoins
		acc.PubKey = pubKey
		ak.SetAccount(ctx, &acc)
		accs = append(accs, &acc)
	}
	return
}
