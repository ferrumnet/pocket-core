package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeeRates(t *testing.T) {
	context, _, keeper := createTestInput(t, true)

	// check initial fee rate
	feeRate := keeper.GetFeeRate(context, "upokt")
	assert.Equal(t, feeRate, uint64(0))

	feeRates := keeper.GetAllFeeRates(context)
	assert.Len(t, feeRates, 0)

	// set fee rate
	err := keeper.SetFeeRate(context, "upokt", 1000)
	assert.Nil(t, err)

	// check fee rate
	feeRate = keeper.GetFeeRate(context, "upokt")
	assert.Equal(t, feeRate, uint64(1000))

	// set fee rate again on same token
	err = keeper.SetFeeRate(context, "upokt", 1100)
	assert.Nil(t, err)

	// check fee rate
	feeRate = keeper.GetFeeRate(context, "upokt")
	assert.Equal(t, feeRate, uint64(1100))

	// check all fee rates
	feeRates = keeper.GetAllFeeRates(context)
	assert.Len(t, feeRates, 1)

	// set fee rate for second token
	err = keeper.SetFeeRate(context, "upokt2", 1000)
	assert.Nil(t, err)

	// check all fee rates
	feeRates = keeper.GetAllFeeRates(context)
	assert.Len(t, feeRates, 2)
}
