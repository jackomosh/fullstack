package services

import (
	"math"

	"github.com/altradits/bursaryhub/backend/models"
)

// FeeCalculator handles fee calculations for the platform
type FeeCalculator struct{}

// CalculateCostBreakdown calculates all fees for a donor deposit
// Returns: conversion fee, gas fee, USDT locked, and disbursement fees
func (f *FeeCalculator) CalculateCostBreakdown(donorDepositUSD float64) *models.CostBreakdownResponse {
	conversionFee := donorDepositUSD * 0.0015 // 0.15% conversion fee
	gasFee := 2.0                            // Fixed gas fee in USD
	usdtLocked := donorDepositUSD - conversionFee - gasFee

	// When disbursed: 1% platform fee, 1% conversion to KSH, 100 KSH withdrawal
	// Assuming 1 USD = 140 KSH approximation
	usdtToKSHRate := 140.0
	
	platformFeeUSD := usdtLocked * 0.01
	platformFeeKSH := platformFeeUSD * usdtToKSHRate
	
	conversionsToKSH := usdtLocked * usdtToKSHRate * 0.01
	withdrawalFee := 100.0
	
	schoolReceivesKSH := (usdtLocked * usdtToKSHRate) - platformFeeKSH - withdrawalFee - conversionsToKSH

	return &models.CostBreakdownResponse{
		DonorDeposit: donorDepositUSD,
		ConversionToUSDT: conversionFee,
		NetworkGasFee: gasFee,
		USDTLocked: usdtLocked,
		WhenDisbursed: models.DisbursementBreakdown{
			USDTAmount: usdtLocked,
			ConversionToKSH: conversionsToKSH,
			WithdrawalFee: withdrawalFee,
			PlatformFee: platformFeeKSH,
			SchoolReceives: schoolReceivesKSH,
		},
		TotalFeesUSD: conversionFee + gasFee,
		TotalFeesKSH: platformFeeKSH + withdrawalFee + conversionsToKSH,
	}
}

// CalculateAutoBalance calculates new balance based on formula:
// New Balance = Previous Unpaid + (New Fee - Paid to Date)
func (f *FeeCalculator) CalculateAutoBalance(previousUnpaid, newFee, paidToDate float64) float64 {
	return previousUnpaid + (newFee - paidToDate)
}

// CalculateTotalFeeForScholarship calculates total fees for a scholarship fund
func (f *FeeCalculator) CalculateTotalFeeForScholarship(amountUSD float64) float64 {
	return math.Round(amountUSD * 0.01 * 140) // 1% in KSH equivalent
}