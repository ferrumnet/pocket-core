package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func (k Keeper) Swap(ctx sdk.Ctx, from string, token string, amount uint64, targetNetwork string,
	targetToken string, targetAddress string) error {
	// TODO: transfer tokens from `from` account to module account by `amount`

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventBridgeSwap,
			sdk.NewAttribute(types.AttributeKeyFrom, from),
			sdk.NewAttribute(types.AttributeKeyToken, token),
			sdk.NewAttribute(types.AttributeKeyTargetNetwork, targetNetwork),
			sdk.NewAttribute(types.AttributeKeyTargetToken, targetToken),
			sdk.NewAttribute(types.AttributeKeyTargetAddress, targetAddress),
			sdk.NewAttribute(types.AttributeKeyAmount, fmt.Sprintf("%d", amount)),
		),
	})

	return nil
}
