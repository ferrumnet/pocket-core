package bridgefee

import (
	"fmt"

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto/keys"
	"github.com/pokt-network/pocket-core/crypto/keys/mintkey"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/auth"
	"github.com/pokt-network/pocket-core/x/auth/util"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/tendermint/tendermint/rpc/client"
)

// SetTokenInfoTx broadcasts a transaction that sends token information for fee distribution
func SetTokenInfoTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, bufferSize uint64, tokenSpecificConfig uint32, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgSetTokenInfo{
		FromAddress: fromAddr,
		Info: types.TokenInfo{
			Token:      token,
			BufferSize: bufferSize,
		},
	}
	txBuilder, cliCtx, err := newTx(cdc, &msg, fromAddr, tmNode, keybase, passphrase)
	if err != nil {
		return nil, err
	}
	err = msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	return util.CompleteAndBroadcastTxCLI(txBuilder, cliCtx, &msg, legacyCodec)
}

// SetTokenTargetInfosTx broadcasts a transaction that sends token targets information for fee distribution
func SetTokenTargetInfosTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, targets []types.TargetInfo, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgSetTokenTargetInfos{
		FromAddress: fromAddr,
		Token:       token,
		Targets:     targets,
	}
	txBuilder, cliCtx, err := newTx(cdc, &msg, fromAddr, tmNode, keybase, passphrase)
	if err != nil {
		return nil, err
	}
	err = msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	return util.CompleteAndBroadcastTxCLI(txBuilder, cliCtx, &msg, legacyCodec)
}

// SetTokenTargetInfosTx broadcasts a transaction that sends token targets information for fee distribution
func SetGlobalTargetInfosTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, targets []types.TargetInfo, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgSetGlobalTargetInfos{
		FromAddress: fromAddr,
		Targets:     targets,
	}
	txBuilder, cliCtx, err := newTx(cdc, &msg, fromAddr, tmNode, keybase, passphrase)
	if err != nil {
		return nil, err
	}
	err = msg.ValidateBasic()
	if err != nil {
		return nil, err
	}
	return util.CompleteAndBroadcastTxCLI(txBuilder, cliCtx, &msg, legacyCodec)
}

// newTx defines generalized transaction broadcast functionality from a proto message
func newTx(cdc *codec.Codec, msg sdk.ProtoMsg, fromAddr sdk.Address, tmNode client.Client, keybase keys.Keybase, passphrase string) (txBuilder auth.TxBuilder, cliCtx util.CLIContext, err error) {
	genDoc, err := tmNode.Genesis()
	if err != nil {
		return
	}
	chainID := genDoc.Genesis.ChainID
	kp, err := keybase.Get(fromAddr)
	if err != nil {
		return
	}
	privkey, err := mintkey.UnarmorDecryptPrivKey(kp.PrivKeyArmor, passphrase)
	if err != nil {
		return
	}
	cliCtx = util.NewCLIContext(tmNode, fromAddr, passphrase).WithCodec(cdc)
	cliCtx.BroadcastMode = util.BroadcastSync
	cliCtx.PrivateKey = privkey
	account, err := cliCtx.GetAccount(fromAddr)
	if err != nil {
		return
	}
	fee := msg.GetFee()
	if account.GetCoins().AmountOf(sdk.DefaultStakeDenom).LT(fee) { // todo get stake denom
		_ = fmt.Errorf("insufficient funds: the fee needed is %v", fee)
		return
	}
	txBuilder = auth.NewTxBuilder(
		auth.DefaultTxEncoder(cdc),
		auth.DefaultTxDecoder(cdc),
		chainID,
		"",
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultStakeDenom, fee))).WithKeybase(keybase)
	return
}
