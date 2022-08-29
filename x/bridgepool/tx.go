package bridgepool

import (
	"fmt"

	"github.com/pokt-network/pocket-core/codec"
	"github.com/pokt-network/pocket-core/crypto/keys"
	"github.com/pokt-network/pocket-core/crypto/keys/mintkey"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/auth"
	"github.com/pokt-network/pocket-core/x/auth/util"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
	"github.com/tendermint/tendermint/rpc/client"
)

// SetFeeTx broadcasts a transaction to set fee rate for a token
func SetFeeTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, fee10000 uint64, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgSetFee{
		FromAddress: fromAddr,
		Token:       token,
		Fee10000:    fee10000,
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

// AllowTargetTx broadcasts a transaction to allow swap target
func AllowTargetTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token, chainId, targetToken string, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgAllowTarget{
		FromAddress: fromAddr,
		Token:       token,
		ChainId:     chainId,
		TargetToken: targetToken,
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

// DisallowTargetTx broadcasts a transaction to disallow swap target
func DisallowTargetTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token, chainId string, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgDisallowTarget{
		FromAddress: fromAddr,
		Token:       token,
		ChainId:     chainId,
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

// AddLiquidityTx broadcasts a transaction to add liquidity into the module
func AddLiquidityTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, amount uint64, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgAddLiquidity{
		FromAddress: fromAddr,
		Token:       token,
		Amount:      amount,
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

// RemoveLiquidityTx broadcasts a transaction to remove liquidity into the module
func RemoveLiquidityTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, amount uint64, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgRemoveLiquidity{
		FromAddress: fromAddr,
		Token:       token,
		Amount:      amount,
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

// SwapTx broadcasts a transaction to swap tokens from current chain to target chain
func SwapTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, token string, amount uint64, targetChainId, targetToken, targetAddress string, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgSwap{
		FromAddress:   fromAddr,
		Token:         token,
		Amount:        amount,
		TargetChainId: targetChainId,
		TargetToken:   targetToken,
		TargetAddress: targetAddress,
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

// WithdrawSignedTx broadcasts a transaction to withdraw tokens from the module based on signer's signature
func WithdrawSignedTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, payee string, amount sdk.Coin, salt string, signature []byte, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgWithdrawSigned{
		FromAddress: fromAddr,
		Payee:       payee,
		Amount:      amount,
		Salt:        salt,
		Signature:   signature,
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

// AddSignerTx broadcasts a transaction to add signer on the module
func AddSignerTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, signer string, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgAddSigner{
		FromAddress: fromAddr,
		Signer:      signer,
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

// AddSignerTx broadcasts a transaction to remove signer from the module
func RemoveSignerTx(cdc *codec.Codec, tmNode client.Client, keybase keys.Keybase, signer string, kp keys.KeyPair, passphrase string, legacyCodec bool) (*sdk.TxResponse, error) {
	fromAddr := kp.GetAddress()
	msg := types.MsgRemoveSigner{
		FromAddress: fromAddr,
		Signer:      signer,
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
