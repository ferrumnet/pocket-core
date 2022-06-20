package keeper

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) Swap(ctx sdk.Ctx, from sdk.Address, token string, amount uint64, targetChainId string,
	targetToken string, targetAddress string) sdk.Error {
	allowedTarget := k.GetAllowedTarget(ctx, token, targetChainId)
	if targetToken != allowedTarget {
		return types.ErrInvalidTargetToken(k.codespace)
	}

	// check ethereum addresses
	if !common.IsHexAddress(token) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}
	if !common.IsHexAddress(targetToken) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}
	if !common.IsHexAddress(targetAddress) {
		return types.ErrInvalidEthereumAddress(k.codespace)
	}

	// transfer tokens from `from` account to module account by `amount`
	err := k.AccountKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, sdk.Coins{sdk.NewInt64Coin(token, int64(amount))})
	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeSwap,
			sdk.NewAttribute(types.AttributeKeyFrom, from.String()),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyTargetChainId, targetChainId),
			sdk.NewAttribute(types.AttributeKeyTargetToken, targetToken),
			sdk.NewAttribute(types.AttributeKeyTargetAddress, targetAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	return nil
}
