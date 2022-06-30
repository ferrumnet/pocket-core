package bridgefee

import (
	"fmt"
	"reflect"

	"github.com/pokt-network/pocket-core/crypto"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/keeper"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
)

func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Ctx, msg sdk.Msg, _ crypto.PublicKey) sdk.Result {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		// convert to value for switch consistency
		if reflect.ValueOf(msg).Kind() == reflect.Ptr {
			msg = reflect.Indirect(reflect.ValueOf(msg)).Interface().(sdk.Msg)
		}

		switch msg := msg.(type) {
		case types.MsgAddAllowedActor:
			return handleMsgAddAllowedActor(ctx, msg, k)
		case types.MsgRemoveAllowedActor:
			return handleMsgRemoveAllowedActor(ctx, msg, k)
		case types.MsgSetTokenInfo:
			return handleMsgSetTokenInfo(ctx, msg, k)
		case types.MsgSetTokenTargetInfos:
			return handleMsgSetTokenTargetInfos(ctx, msg, k)
		case types.MsgSetGlobalTargetInfos:
			return handleMsgSetGlobalTargetInfos(ctx, msg, k)
		default:
			errMsg := fmt.Sprintf("unrecognized bridgefee message type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgAddAllowedActor(ctx sdk.Ctx, msg types.MsgAddAllowedActor, k keeper.Keeper) sdk.Result {
	moduleOwner := k.GetParams(ctx).Owner
	if msg.FromAddress.String() != moduleOwner {
		return types.ErrNotEnoughPermission(k.Codespace()).Result()
	}

	err := k.AddAllowedActor(ctx, msg.Actor)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeAddAllowedActor,
			sdk.NewAttribute(types.AttributeKeyActor, msg.Actor),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgRemoveAllowedActor(ctx sdk.Ctx, msg types.MsgRemoveAllowedActor, k keeper.Keeper) sdk.Result {
	moduleOwner := k.GetParams(ctx).Owner
	if msg.FromAddress.String() != moduleOwner {
		return types.ErrNotEnoughPermission(k.Codespace()).Result()
	}

	err := k.RemoveAllowedActor(ctx, msg.Actor)
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRemoveAllowedActor,
			sdk.NewAttribute(types.AttributeKeyActor, msg.Actor),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSetTokenInfo(ctx sdk.Ctx, msg types.MsgSetTokenInfo, k keeper.Keeper) sdk.Result {
	moduleOwner := k.GetParams(ctx).Owner
	if msg.FromAddress.String() != moduleOwner {
		return types.ErrNotEnoughPermission(k.Codespace()).Result()
	}

	err := k.SetTokenInfo(ctx, msg.Info)
	if err != nil {
		return err.Result()
	}

	// ctx.EventManager().EmitEvents(sdk.Events{
	// 	sdk.NewEvent(
	// 		types.EventTypeSetSigner,
	// 		sdk.NewAttribute(types.AttributeKeySigner, msg.Signer),
	// 	)},
	// )
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSetTokenTargetInfos(ctx sdk.Ctx, msg types.MsgSetTokenTargetInfos, k keeper.Keeper) sdk.Result {
	moduleOwner := k.GetParams(ctx).Owner
	if msg.FromAddress.String() != moduleOwner {
		return types.ErrNotEnoughPermission(k.Codespace()).Result()
	}

	err := k.SetTokenTargetInfo(ctx, types.TokenTargetInfo{
		Token:   msg.Token,
		Targets: msg.Targets,
	})
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetTokenTargetInfos,
			sdk.NewAttribute(types.AttributeKeyToken, msg.Token),
			sdk.NewAttribute(types.AttributeKeySender, msg.FromAddress.String()),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}

func handleMsgSetGlobalTargetInfos(ctx sdk.Ctx, msg types.MsgSetGlobalTargetInfos, k keeper.Keeper) sdk.Result {
	moduleOwner := k.GetParams(ctx).Owner
	if msg.FromAddress.String() != moduleOwner {
		return types.ErrNotEnoughPermission(k.Codespace()).Result()
	}

	err := k.SetTokenTargetInfo(ctx, types.TokenTargetInfo{
		Token:   "",
		Targets: msg.Targets,
	})
	if err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSetGlobalTokenTargetInfos,
			sdk.NewAttribute(types.AttributeKeySender, msg.FromAddress.String()),
		)},
	)
	return sdk.Result{Events: ctx.EventManager().Events()}
}
