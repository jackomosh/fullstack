package services

import (
	"fmt"
	"os"
	"time"
)

// M-Pesa service handles STK push payments
type MpesService struct {
	consumerKey    string
	consumerSecret string
	shortcode      string
}

// NewMpesService creates a new M-Pesa service instance
func NewMpesService() *MpesService {
	return &MpesService{
		consumerKey:    os.Getenv("MPESA_CONSUMER_KEY"),
		consumerSecret: os.Getenv("MPESA_CONSUMER_SECRET"),
		shortcode:      os.Getenv("MPESA_SHORTCODE"),
	}
}

// STKPush initiates STK push payment to donor
func (m *MpesService) STKPush(phone, amount float64, accountRef string) (string, error) {
	if m.consumerKey == "" {
		// Stub mode
		fmt.Printf("[STUB] STK Push to %s for %f KSH, ref: %s\n", phone, amount, accountRef)
		return "stub-tx-" + fmt.Sprintf("%d", time.Now().Unix()), nil
	}
	// In production, integrate with M-Pesa API
	return "", nil
}

// Callback handles M-Pesa payment callback
func (m *MpesService) Callback(data map[string]interface{}) error {
	// In production, verify and update payment status
	fmt.Printf("[STUB] M-Pesa callback: %+v\n", data)
	return nil
}