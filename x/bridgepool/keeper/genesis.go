package keeper

import (
	sdk "github.com/pokt-network/pocket-core/types"
	"github.com/pokt-network/pocket-core/x/bridgepool/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis - Init store state from genesis data
func (k Keeper) InitGenesis(ctx sdk.Ctx, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)
	for _, signer := range data.Signers {
		k.SetSigner(ctx, signer)
	}
	for _, liq := range data.Liquidities {
		addr, err := sdk.AddressFromHex(liq.Address)
		if err != nil {
			panic(err)
		}
		k.SetLiquidity(ctx, liq.Token, addr, liq.Amount)
	}
	for _, rate := range data.Fees {
		k.SetFeeRate(ctx, rate.Token, rate.Rate)
	}
	for _, target := range data.AllowedTargets {
		k.AllowTarget(ctx, target.Token, target.ChainId, target.TargetToken)
	}
	for _, salt := range data.UsedWithdrawSalts {
		k.SetUsedSalt(ctx, salt)
	}
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns a GenesisState for a given context and keeper
func (k Keeper) ExportGenesis(ctx sdk.Ctx) types.GenesisState {
	return types.NewGenesisState(
		k.GetParams(ctx),
		k.GetAllSigners(ctx),
		k.GetAllLiquidities(ctx),
		k.GetAllFeeRates(ctx),
		k.GetAllAllowedTargets(ctx),
		k.GetAllUsedSalts(ctx),
	)
}
