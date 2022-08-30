package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
)

func (k Keeper) DistributeTax(ctx sdk.Ctx, token string) {
	tokenInfo := k.GetTokenInfo(ctx, token)
	moduleAddr := k.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	moduleCoins := k.AccountKeeper.GetCoins(ctx, moduleAddr.GetAddress())
	moduleToken := moduleCoins.AmountOf(token)

	if moduleToken.LT(sdk.NewInt(int64(tokenInfo.BufferSize))) {
		return
	}

	tokenTargetConfig := k.GetTokenTargetInfo(ctx, token)
	if tokenTargetConfig.Token == "" { // if token target config does not exist
		globalTargetConfig := k.GetGlobalTokenTargetInfo(ctx)
		tokenTargetConfig = globalTargetConfig
	}

	totalWeight := uint64(0)

	for _, target := range tokenTargetConfig.Targets {
		totalWeight += target.Weight
	}

	if totalWeight == 0 {
		return
	}
	for _, target := range tokenTargetConfig.Targets {
		targetAmount := moduleToken.Mul(sdk.NewInt(int64(target.Weight))).Quo(sdk.NewInt(int64(totalWeight)))
		cacheCtx, write := ctx.CacheContext()
		if targetAmount.IsPositive() {
			switch target.TType {
			case types.TargetType_Burn:
				err := k.AccountKeeper.BurnCoins(cacheCtx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(token, targetAmount)))
				if err != nil {
					fmt.Println(err)
				} else {
					write()
				}
			case types.TargetType_Address:
				addr, err := sdk.AddressFromHex(target.Target)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = k.AccountKeeper.SendCoinsFromModuleToAccount(cacheCtx, types.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(token, targetAmount)))
				if err != nil {
					fmt.Println(err)
				} else {
					write()
				}
			}
		}
	}
}
