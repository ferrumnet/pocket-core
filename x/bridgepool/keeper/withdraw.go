package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) WithdrawSigned(ctx sdk.Ctx, from string, token string, payee string, amount uint64,
	salt []byte, signature []byte) sdk.Error {
	// TODO: verify signature
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
	//     bytes32 message = withdrawSignedMessage(token, payee, amount, salt);
	//     address _signer = signerUnique(message, signature);
	signer := ""
	//     require(signers[_signer], "BridgePool: Invalid signer");

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
			sdk.NewAttribute(types.AttributeKeySigner, signer),
			sdk.NewAttribute(types.AttributeKeyReceiver, payee),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
			sdk.NewAttribute(types.AttributeKeyFee, fmt.Sprintf("%d", fee)),
		),
	})

	return nil
}
