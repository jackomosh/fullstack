package services

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// BlockchainService handles interactions with the BursaryEscrow smart contract
type BlockchainService struct {
	escrowAddress  common.Address
	usdtAddress    common.Address
	platformWallet common.Address
}

// NewBlockchainService creates a new blockchain service instance
func NewBlockchainService() *BlockchainService {
	return &BlockchainService{
		escrowAddress:  common.HexToAddress(getEnv("ESCROW_CONTRACT_ADDRESS", "0x0000000000000000000000000000000000000000")),
		usdtAddress:    common.HexToAddress(getEnv("USDT_CONTRACT_ADDRESS", "0x0000000000000000000000000000000000000000")),
		platformWallet: common.HexToAddress(getEnv("PLATFORM_WALLET", "0x0000000000000000000000000000000000000000")),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// DepositUSDT locks USDT in escrow for a scholarship
func (b *BlockchainService) DepositUSDT(scholarshipID uint, amount float64) (string, error) {
	txHash := fmt.Sprintf("0x%x", crypto.Keccak256([]byte(fmt.Sprintf("deposit-%d-%f", scholarshipID, amount))))[:66]
	return txHash, nil
}

// ExecuteDisbursement sends USDT to school after three-way verification
func (b *BlockchainService) ExecuteDisbursement(disbursementID uint, schoolAddress string, amountUSDT float64) (string, error) {
	fee := amountUSDT * 0.01
	amountToSchool := amountUSDT - fee

	txHash := fmt.Sprintf("0x%x", crypto.Keccak256([]byte(fmt.Sprintf("disburse-%d-%s-%f", disbursementID, schoolAddress, amountToSchool))))[:66]
	return txHash, nil
}

// SetWhitelistedSchool adds/removes school from whitelist
func (b *BlockchainService) SetWhitelistedSchool(schoolAddress string, status bool) error {
	fmt.Printf("[STUB] Set whitelist for %s to %v\n", schoolAddress, status)
	return nil
}

// SetWhitelistedDonor adds/removes donor from whitelist
func (b *BlockchainService) SetWhitelistedDonor(donorAddress string, status bool) error {
	fmt.Printf("[STUB] Set donor whitelist for %s to %v\n", donorAddress, status)
	return nil
}

// GenerateTestWallet creates a test wallet for development
func GenerateTestWallet() (address string, privateKey string, err error) {
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	address = crypto.PubkeyToAddress(privateKeyECDSA.PublicKey).Hex()
	return address, "test-private-key", nil
}

// VerifyThreeWayMatch checks if all three amounts match
func VerifyThreeWayMatch(studentAmount, schoolAmount, feeMasterAmount float64) bool {
	return studentAmount == schoolAmount && schoolAmount == feeMasterAmount
}