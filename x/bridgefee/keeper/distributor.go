package keeper

import (
	"fmt"

	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgefee/types"
)

// DistributeTax implements distribution of specific token
func (k Keeper) DistributeTax(ctx sdk.Ctx, token string) {
	// get token info to be used for token distribution rule
	tokenInfo := k.GetTokenInfo(ctx, token)

	// get module account balances stacked
	moduleAddr := k.AccountKeeper.GetModuleAccount(ctx, types.ModuleName)
	moduleCoins := k.AccountKeeper.GetCoins(ctx, moduleAddr.GetAddress())
	moduleToken := moduleCoins.AmountOf(token)

	// check buffer amount passed for the token
	if moduleToken.LT(sdk.NewInt(int64(tokenInfo.BufferSize))) {
		return
	}

	// get token target information to distribute to
	tokenTargetConfig := k.GetTokenTargetInfo(ctx, token)
	if tokenTargetConfig.Token == "" { // if token target config does not exist
		globalTargetConfig := k.GetGlobalTokenTargetInfo(ctx)
		tokenTargetConfig = globalTargetConfig
	}

	// get total weights of targets to do weighted distribution
	totalWeight := uint64(0)
	for _, target := range tokenTargetConfig.Targets {
		totalWeight += target.Weight
	}

	// distribution logic implementation
	if totalWeight == 0 {
		return
	}
	for _, target := range tokenTargetConfig.Targets {
		// calculate distribution amount based on weight
		targetAmount := moduleToken.Mul(sdk.NewInt(int64(target.Weight))).Quo(sdk.NewInt(int64(totalWeight)))
		cacheCtx, write := ctx.CacheContext()
		// distribute tokens based on distribution type
		if targetAmount.IsPositive() {
			switch target.TType {
			case types.TargetType_Burn: // burn coins if target type is burn
				err := k.AccountKeeper.BurnCoins(cacheCtx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(token, targetAmount)))
				if err != nil {
					fmt.Println(err)
				} else {
					write()
				}
			case types.TargetType_Address: // send tokens to an address if target type is an address
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
