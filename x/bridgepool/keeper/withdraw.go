package keeper

import (
	"bytes"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	bridgefeeTypes "github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) SetUsedMessage(ctx sdk.Ctx, salt []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.WithdrawSaltKey(salt), salt)
}

func (k Keeper) IsUsedMessage(ctx sdk.Ctx, message []byte) bool {
	store := ctx.KVStore(k.storeKey)
	bz, _ := store.Get(types.WithdrawSaltKey(message))
	return bytes.Equal(bz, message)
}

func (k Keeper) GetAllUsedMessages(ctx sdk.Ctx) [][]byte {
	usedMessages := [][]byte{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.WithdrawSaltKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		usedMessages = append(usedMessages, iterator.Value())
	}
	return usedMessages
}

func GetSigner(chainId string, payee string, amount sdk.Coin,
	salt string, signature []byte) (common.Address, []byte, error) {
	signer := common.Address{}

	// verify signature
	message := &types.WithdrawSignMessage{
		ChainId: chainId,
		Payee:   payee,
		Amount:  amount,
		Salt:    salt,
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return signer, messageBytes, err
	}
	if len(signature) > crypto.RecoveryIDOffset {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
		recovered, err := crypto.SigToPub(accounts.TextHash(messageBytes), signature)
		if err != nil {
			return signer, messageBytes, err
		}
		signer = crypto.PubkeyToAddress(*recovered)
	}

	return signer, messageBytes, nil
}

func (k Keeper) WithdrawSigned(ctx sdk.Ctx, from string, payee string, amount sdk.Coin,
	salt string, signature []byte) sdk.Error {

	// verify signature
	signer, messageBytes, err := GetSigner(ctx.ChainID(), payee, amount, salt, signature)
	if err != nil {
		return types.ErrUnexpectedError(k.codespace, err)
	}

	// TODO: enable this when goes live
	// if !k.IsSigner(ctx, signer.String()) {
	// 	return types.ErrInvalidSigner(k.codespace)
	// }

	// avoid using same signature and salt again
	if k.IsUsedMessage(ctx, messageBytes) {
		return types.ErrAlreadyUsedWithdrawMessage(k.codespace)
	}

	// handle fees
	feeRate := k.GetFeeRate(ctx, amount.Denom)
	fee := amount.Amount.Mul(sdk.NewInt(int64(feeRate))).Quo(sdk.NewInt(int64(10000)))
	amountWithoutFee := amount.Amount
	if fee.IsPositive() {
		err := k.AccountKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, bridgefeeTypes.ModuleName, sdk.Coins{sdk.NewCoin(amount.Denom, fee)})
		if err != nil {
			return types.ErrUnexpectedError(k.codespace, err)
		}

		k.bridgeFeeKeeper.DistributeTax(ctx, amount.Denom)
		amountWithoutFee = amountWithoutFee.Sub(fee)
	}

	// transfer amount except fee to payee account
	payeeAcc, err := sdk.AddressFromHex(payee)
	err = k.AccountKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, payeeAcc, sdk.Coins{sdk.NewCoin(amount.Denom, amountWithoutFee)})
	if err != nil {
		return types.ErrUnexpectedError(k.codespace, err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTransferBySignature,
			sdk.NewAttribute(types.AttributeKeySigner, signer.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, payee),
			sdk.NewAttribute(types.AttributeKeyToken, amount.Denom),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyFee, fee.String()),
		),
	})

	return nil
}
