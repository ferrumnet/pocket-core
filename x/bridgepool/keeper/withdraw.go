package keeper

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func withdrawSignedMessage() []byte {
	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	arguments := abi.Arguments{
		{
			Type: addressTy,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: uint256Ty,
		},
	}

	bytes, _ := arguments.Pack(
		common.HexToAddress("0x0000000000000000000000000000000000000000"),
		[32]byte{'I', 'D', '1'},
		big.NewInt(42),
	)

	var buf []byte
	hash := sha3.NewKeccak256()
	hash.Write(bytes)
	buf = hash.Sum(buf)

	fmt.Println(hexutil.Encode(buf))
	return buf
}

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
