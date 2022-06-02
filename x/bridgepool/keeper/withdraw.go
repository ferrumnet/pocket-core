package keeper

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	sdk "github.com/pokt-network/pocket-core/types"
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

func (k Keeper) WithdrawSigned(ctx sdk.Ctx, from string, token string, payee string, amount uint64,
	salt []byte, signature []byte) sdk.Error {
	// verify signature
	message := withdrawSignedMessage(token, payee, amount, salt)

	signature[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	recovered, err := crypto.SigToPub(message.Bytes(), signature)
	if err != nil {
		return types.ErrInvalidSignature(k.codespace, err)
	}
	signer := crypto.PubkeyToAddress(*recovered)
	// TODO: enable this when goes live
	// if !k.IsSigner(ctx, signer.String()) {
	// 	return types.ErrInvalidSigner(k.codespace)
	// }

	// TODO: avoid using same signature and salt again

	// TODO: handle fees
	feeRate := k.GetFeeRate(ctx, token)
	fee := amount * feeRate / 10000
	if fee != 0 {
		// TODO: transfer fee amount to fee handler account
		// TODO: distribute fees by fee handler
		amount -= fee
	}

	// TODO: transfer remaining amount to the payee

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
