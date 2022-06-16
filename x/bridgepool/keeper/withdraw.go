package keeper

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
	bridgefeeTypes "github.com/pokt-network/pocket-core/x/bridgefee/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func withdrawSignedMessage(token string, payee string, amount uint64, salt []byte) common.Hash {
	// function withdrawSignedMessage(
	//         address token,
	//         address payee,
	//         uint256 amount,
	//         bytes32 salt)
	// internal pure returns (bytes32) {
	//     return keccak256(abi.encode(
	//       WITHDRAW_SIGNED_METHOD,
	//       token,
	//       payee,
	//       amount,
	//       salt
	//     ));
	// }

	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	arguments := abi.Arguments{
		{
			Type: bytes32Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: bytes32Ty,
		},
	}

	WITHDRAW_SIGNED_METHOD := crypto.Keccak256Hash([]byte("WithdrawSigned(address token,address payee,uint256 amount,bytes32 salt)"))
	bytes, _ := arguments.Pack(
		WITHDRAW_SIGNED_METHOD,
		token,
		payee,
		amount,
		salt,
	)

	return crypto.Keccak256Hash(bytes)
}

func (k Keeper) SetUsedSalt(ctx sdk.Ctx, salt []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.WithdrawSaltKey(salt), salt)
}

func (k Keeper) IsUsedSalt(ctx sdk.Ctx, salt []byte) bool {
	store := ctx.KVStore(k.storeKey)
	bz, _ := store.Get(types.WithdrawSaltKey(salt))
	return bytes.Equal(bz, salt)
}

func (k Keeper) GetAllUsedSalts(ctx sdk.Ctx) [][]byte {
	usedSalts := [][]byte{}
	store := ctx.KVStore(k.storeKey)
	iterator, _ := sdk.KVStorePrefixIterator(store, types.WithdrawSaltKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		usedSalts = append(usedSalts, iterator.Value())
	}
	return usedSalts
}

func (k Keeper) WithdrawSigned(ctx sdk.Ctx, from string, token string, payee string, amount uint64,
	salt []byte, signature []byte) sdk.Error {

	// check ethereum addresses
	if !common.IsHexAddress(token) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}
	if !common.IsHexAddress(token) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}

	// verify signature
	message := withdrawSignedMessage(token, payee, amount, salt)

	signer := common.Address{}
	if len(signature) > crypto.RecoveryIDOffset {
		signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
		recovered, err := crypto.SigToPub(message.Bytes(), signature)
		if err != nil {
			return types.ErrInvalidSignature(k.codespace, err)
		}
		signer = crypto.PubkeyToAddress(*recovered)
	}

	// TODO: enable this when goes live
	// if !k.IsSigner(ctx, signer.String()) {
	// 	return types.ErrInvalidSigner(k.codespace)
	// }

	// avoid using same signature and salt again
	if k.IsUsedSalt(ctx, salt) {
		return types.ErrAlreadyUsedWithdrawSalt(k.codespace)
	}

	// handle fees
	feeRate := k.GetFeeRate(ctx, token)
	fee := amount * feeRate / 10000
	if fee != 0 {
		err := k.AccountKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, bridgefeeTypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(token, int64(fee))})
		if err != nil {
			return types.ErrUnexpectedError(k.codespace, err)
		}

		amount -= fee
	}

	// transfer amount except fee to payee account
	payeeAcc, err := sdk.AddressFromHex(payee)
	err = k.AccountKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, payeeAcc, sdk.Coins{sdk.NewInt64Coin(token, int64(amount))})
	if err != nil {
		return types.ErrUnexpectedError(k.codespace, err)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTransferBySignature,
			sdk.NewAttribute(types.AttributeKeySigner, signer.String()),
			sdk.NewAttribute(types.AttributeKeyReceiver, payee),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
			sdk.NewAttribute(types.AttributeKeyFee, fmt.Sprintf("%d", fee)),
		),
	})

	return nil
}
