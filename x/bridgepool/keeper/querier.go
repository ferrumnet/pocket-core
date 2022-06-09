package keeper

import (
	"fmt"

	"github.com/pokt-network/pocket-core/codec"
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// creates a querier for bridgepool REST endpoints
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Ctx, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case "bridgepoolParams":
			return queryParams(ctx, req, k)
		case "bridgepoolAllSigners":
			return queryAllSigners(ctx, req, k)
		case "bridgepoolAllLiquidities":
			return queryAllLiquidities(ctx, req, k)
		case "bridgepoolAllFeeRates":
			return queryAllFeeRates(ctx, req, k)
		case "bridgepoolAllAllowedTargets":
			return queryAllAllowedTargets(ctx, req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown bridgepool query endpoint")
		}
	}
}

func queryParams(ctx sdk.Ctx, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}
	moduleParams := k.GetParams(ctx)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, moduleParams)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}
	return res, nil
}

func queryAllSigners(ctx sdk.Ctx, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}
	allSigners := k.GetAllSigners(ctx)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, allSigners)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}
	return res, nil
}

func queryAllLiquidities(ctx sdk.Ctx, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}
	allLiquidities := k.GetAllLiquidities(ctx)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, allLiquidities)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}
	return res, nil
}

func queryAllFeeRates(ctx sdk.Ctx, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}
	allFeeRates := k.GetAllFeeRates(ctx)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, allFeeRates)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}
	return res, nil
}

func queryAllAllowedTargets(ctx sdk.Ctx, req abci.RequestQuery, k Keeper) ([]byte, sdk.Error) {
	var params types.QueryParams
	err := types.ModuleCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}
	allAllowedTargets := k.GetAllAllowedTargets(ctx)
	res, err := codec.MarshalJSONIndent(types.ModuleCdc, allAllowedTargets)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}
	return res, nil
}
