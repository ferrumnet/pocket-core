package keeper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeeRates(t *testing.T) {
	context, _, keeper, _ := createTestInput(t, true)

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

func TestAllowedTarget(t *testing.T) {
	context, _, keeper, _ := createTestInput(t, true)

	// check initial allowed target
	target := keeper.GetAllowedTarget(context, "upokt", "137")
	assert.Equal(t, target, "")

	targets := keeper.GetAllAllowedTargets(context)
	assert.Len(t, targets, 0)

	// set allowed target
	err := keeper.AllowTarget(context, "upokt", "137", "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// check allowed target
	target = keeper.GetAllowedTarget(context, "upokt", "137")
	assert.Equal(t, target, "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")

	// set allowed target again on same token
	err = keeper.AllowTarget(context, "upokt", "137", "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// check fee rate
	target = keeper.GetAllowedTarget(context, "upokt", "137")
	assert.Equal(t, target, "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")

	// check all allowed targets
	targets = keeper.GetAllAllowedTargets(context)
	assert.Len(t, targets, 1)

	// set allowed target for second chain
	err = keeper.AllowTarget(context, "upokt", "13", "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// check all fee rates
	targets = keeper.GetAllAllowedTargets(context)
	assert.Len(t, targets, 2)

	// disallow target
	keeper.DisallowTarget(context, "upokt", "13")

	// check after disallow
	target = keeper.GetAllowedTarget(context, "upokt", "13")
	assert.Equal(t, target, "")

	// check all fee rates
	targets = keeper.GetAllAllowedTargets(context)
	assert.Len(t, targets, 1)
}

func TestSigners(t *testing.T) {
	context, _, keeper, _ := createTestInput(t, true)

	// check if signer when not set
	isSigner := keeper.IsSigner(context, "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Equal(t, isSigner, false)

	signers := keeper.GetAllSigners(context)
	assert.Len(t, signers, 0)

	// set signer
	err := keeper.SetSigner(context, "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// check if signer when it is set
	isSigner = keeper.IsSigner(context, "0x7B848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Equal(t, isSigner, true)

	// check all signers
	signers = keeper.GetAllSigners(context)
	assert.Len(t, signers, 1)

	// set second signer
	err = keeper.SetSigner(context, "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Nil(t, err)

	// check all fee rates
	signers = keeper.GetAllSigners(context)
	assert.Len(t, signers, 2)

	// remove signer
	keeper.DeleteSigner(context, "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")

	// check if signer after removal
	isSigner = keeper.IsSigner(context, "0x8A848510E92B2f2F7ea06d46e7B370198F7369Bc")
	assert.Equal(t, isSigner, false)

	// check all signers
	signers = keeper.GetAllSigners(context)
	assert.Len(t, signers, 1)
}
