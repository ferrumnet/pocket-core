package types

// NewGenesisState - Create a new genesis state
func NewGenesisState(params Params, signers []string, liquidities []TokenLiquidity, fees []string) GenesisState {
	return GenesisState{
		Params:      params,
		Signers:     signers,
		Liquidities: liquidities,
		Fees:        fees,
	}
}

// DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]string{},
		[]TokenLiquidity{},
		[]string{},
	)
}

// ValidateGenesis performs basic validation of auth genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	return nil
}
