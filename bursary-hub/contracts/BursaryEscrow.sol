// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract BursaryEscrow is Ownable, ReentrancyGuard {
    IERC20 public usdt;
    
    struct Disbursement {
        uint256 amount;
        address student;
        address school;
        bool completed;
        uint256 timestamp;
    }
    
    mapping(uint256 => Disbursement) public disbursements;
    mapping(address => bool) public whitelistedSchools;
    mapping(address => bool) public whitelistedDonors;
    
    uint256 public platformFeeBps = 100; // 1% = 100 basis points
    address public feeCollector;
    uint256 public nextDisbursementId = 1;
    
    event FundsDeposited(address indexed donor, uint256 amount, uint256 scholarshipId);
    event DisbursementExecuted(uint256 indexed disbursementId, address school, uint256 amountUSDT, uint256 amountKSH);
    event FeeCollected(uint256 amount, uint256 fee);
    event SchoolWhitelisted(address indexed school, bool status);
    event DonorWhitelisted(address indexed donor, bool status);
    
    constructor(address _usdt) {
        usdt = IERC20(_usdt);
        feeCollector = msg.sender;
    }
    
    function deposit(uint256 scholarshipId, uint256 amount) external nonReentrant {
        require(whitelistedDonors[msg.sender], "Not whitelisted");
        require(usdt.transferFrom(msg.sender, address(this), amount), "Transfer failed");
        emit FundsDeposited(msg.sender, amount, scholarshipId);
    }
    
    function executeDisbursement(
        uint256 disbursementId,
        address school,
        uint256 amountUSDT,
        bytes32 threeWayMatchHash
    ) external onlyOwner nonReentrant {
        require(whitelistedSchools[school], "School not whitelisted");
        require(!disbursements[disbursementId].completed, "Already completed");
        
        uint256 fee = (amountUSDT * platformFeeBps) / 10000;
        uint256 amountToSchool = amountUSDT - fee;
        
        require(usdt.transfer(feeCollector, fee), "Fee transfer failed");
        require(usdt.transfer(school, amountToSchool), "School transfer failed");
        
        disbursements[disbursementId] = Disbursement({
            amount: amountToSchool,
            student: address(0),
            school: school,
            completed: true,
            timestamp: block.timestamp
        });
        
        emit DisbursementExecuted(disbursementId, school, amountToSchool, amountUSDT);
        emit FeeCollected(amountUSDT, fee);
    }
    
    // Soulbound: Prevent direct token transfers
    function transfer(address, uint256) external pure returns (bool) {
        revert("BursaryFund: Tokens are Soulbound");
    }
    
    function setWhitelistedSchool(address school, bool status) external onlyOwner {
        whitelistedSchools[school] = status;
        emit SchoolWhitelisted(school, status);
    }
    
    function setWhitelistedDonor(address donor, bool status) external onlyOwner {
        whitelistedDonors[donor] = status;
        emit DonorWhitelisted(donor, status);
    }
    
    function setFeeCollector(address _feeCollector) external onlyOwner {
        feeCollector = _feeCollector;
    }
}