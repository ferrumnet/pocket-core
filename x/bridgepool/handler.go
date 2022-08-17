package bridgepool

import (
	"fmt"
	"reflect"

	"github.com/pokt-network/pocket-core/crypto"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/keeper"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Ctx, msg sdk.Msg, _ crypto.PublicKey) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		// convert to value for switch consistency
		if reflect.ValueOf(msg).Kind() == reflect.Ptr {
			msg = reflect.Indirect(reflect.ValueOf(msg)).Interface().(sdk.Msg)
		}

		switch msg := msg.(type) {
		case types.MsgSetFee:
			return handleMsgSetFee(ctx, msg, k)
		case types.MsgAllowTarget:
			return handleMsgAllowTarget(ctx, msg, k)
		case types.MsgDisallowTarget:
			return handleMsgDisallowTarget(ctx, msg, k)
		case types.MsgAddLiquidity:
			return handleMsgAddLiquidity(ctx, msg, k)
		case types.MsgRemoveLiquidity:
			return handleMsgRemoveLiquidity(ctx, msg, k)
		case types.MsgSwap:
			return handleMsgSwap(ctx, msg, k)
		case types.MsgWithdrawSigned:
			return handleMsgWithdrawSigned(ctx, msg, k)
		case types.MsgAddSigner:
			return handleMsgAddSigner(ctx, msg, k)
		case types.MsgRemoveSigner:
			return handleMsgRemoveSigner(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized bridgepool message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddSigner(ctx sdk.Ctx, msg types.MsgAddSigner, k keeper.Keeper) sdk.Result {
	// TODO: enable this when goes live
	// moduleOwner := k.GetParams(ctx).Owner
	// if msg.FromAddress.String() != moduleOwner {
	// 	return types.ErrNotEnoughPermission(k.Codespace()).Result()
	// }

	err := k.SetSigner(ctx, msg.Signer)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetSigner,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Signer),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgRemoveSigner(ctx sdk.Ctx, msg types.MsgRemoveSigner, k keeper.Keeper) sdk.Result {
	// TODO: enable this when goes live
	// moduleOwner := k.GetParams(ctx).Owner
	// if msg.FromAddress.String() != moduleOwner {
	// 	return types.ErrNotEnoughPermission(k.Codespace()).Result()
	// }

	err := k.DeleteSigner(ctx, msg.Signer)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveSigner,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Signer),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSetFee(ctx sdk.Ctx, msg types.MsgSetFee, k keeper.Keeper) sdk.Result {
	// TODO: enable this when goes live
	// moduleOwner := k.GetParams(ctx).Owner
	// if msg.FromAddress.String() != moduleOwner {
	// 	return types.ErrNotEnoughPermission(k.Codespace()).Result()
	// }

	err := k.SetFeeRate(ctx, msg.Token, msg.Fee10000)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetFeeRate,
			sdk.NewAttribute(types.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types.AttributeKeyFee, fmt.Sprintf("%d", msg.Fee10000)),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgAllowTarget(ctx sdk.Ctx, msg types.MsgAllowTarget, k keeper.Keeper) sdk.Result {
	// TODO: enable this when goes live
	// moduleOwner := k.GetParams(ctx).Owner
	// if msg.FromAddress.String() != moduleOwner {
	// 	return types.ErrNotEnoughPermission(k.Codespace()).Result()
	// }

	err := k.AllowTarget(ctx, msg.Token, msg.ChainId, msg.TargetToken)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAllowTarget,
			sdk.NewAttribute(types.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types.AttributeKeyChainId, msg.ChainId),
			sdk.NewAttribute(types.AttributeKeyTargetToken, msg.TargetToken),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgDisallowTarget(ctx sdk.Ctx, msg types.MsgDisallowTarget, k keeper.Keeper) sdk.Result {
	// TODO: enable this when goes live
	// moduleOwner := k.GetParams(ctx).Owner
	// if msg.FromAddress.String() != moduleOwner {
	// 	return types.ErrNotEnoughPermission(k.Codespace()).Result()
	// }

	k.DisallowTarget(ctx, msg.Token, msg.ChainId)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDisallowTarget,
			sdk.NewAttribute(types.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types.AttributeKeyChainId, msg.ChainId),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgAddLiquidity(ctx sdk.Ctx, msg types.MsgAddLiquidity, k keeper.Keeper) sdk.Result {
	err := k.AddLiquidity(ctx, msg.Token, msg.FromAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgRemoveLiquidity(ctx sdk.Ctx, msg types.MsgRemoveLiquidity, k keeper.Keeper) sdk.Result {
	err := k.RemoveLiquidity(ctx, msg.Token, msg.FromAddress, msg.Amount)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSwap(ctx sdk.Ctx, msg types.MsgSwap, k keeper.Keeper) sdk.Result {
	err := k.Swap(ctx, msg.FromAddress, msg.Token, msg.Amount, msg.TargetChainId, msg.TargetToken, msg.TargetAddress)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgWithdrawSigned(ctx sdk.Ctx, msg types.MsgWithdrawSigned, k keeper.Keeper) sdk.Result {
	err := k.WithdrawSigned(ctx, msg.FromAddress.String(), msg.Payee, msg.Amount, msg.Salt, msg.Signature)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{Events: ctx.EventManager().Events()}
}
