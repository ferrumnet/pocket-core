package bridgepool

import (
	"fmt"

	"github.com/pokt-network/pocket-core/crypto"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/keeper"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Ctx, msg sdk.Msg, _ crypto.PublicKey) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		default:
			errMsg := fmt.Sprintf("unrecognized bridgepool message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
