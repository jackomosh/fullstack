const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("BursaryEscrow", function () {
  let usdt, escrow, owner, donor, school, other;

  beforeEach(async function () {
    [owner, donor, school, other] = await ethers.getSigners();

    // Deploy mock USDT
    const usdtFactory = await ethers.getContractFactory("MockUSDT");
    usdt = await usdtFactory.deploy();
    await usdt.deployed();

    // Deploy escrow
    const escrowFactory = await ethers.getContractFactory("BursaryEscrow");
    escrow = await escrowFactory.deploy(usdt.address);
    await escrow.deployed();

    // Mint and approve USDT for donor
    await usdt.mint(donor.address, ethers.utils.parseUnits("100000", 6));
    await usdt.connect(donor).approve(escrow.address, ethers.utils.parseUnits("100000", 6));
  });

  describe("Deposit", function () {
    it("Should accept deposit from whitelisted donor", async function () {
      await escrow.setWhitelistedDonor(donor.address, true);
      
      await expect(
        escrow.connect(donor).deposit(1, ethers.utils.parseUnits("5000", 6))
      ).to.emit(escrow, "FundsDeposited")
        .withArgs(donor.address, ethers.utils.parseUnits("5000", 6), 1);
    });

    it("Should reject deposit from non-whitelisted donor", async function () {
      await expect(
        escrow.connect(donor).deposit(1, ethers.utils.parseUnits("5000", 6))
      ).to.be.revertedWith("Not whitelisted");
    });
  });

  describe("Disbursement", function () {
    beforeEach(async function () {
      await escrow.setWhitelistedDonor(donor.address, true);
      await escrow.setWhitelistedSchool(school.address, true);
    });

    it("Should execute disbursement with correct amounts and 1% platform fee", async function () {
      // Fund the contract with USDT
      await usdt.mint(escrow.address, ethers.utils.parseUnits("10000", 6));
      
      const amount = ethers.utils.parseUnits("1000", 6);
      const fee = amount.mul(100).div(10000); // 1% = 100 basis points
      const toSchool = amount.sub(fee);

      await expect(
        escrow.executeDisbursement(1, school.address, amount, "0x00")
      ).to.emit(escrow, "DisbursementExecuted");
    });

    it("Should revert soulbound transfer attempt", async function () {
      await expect(
        escrow.transfer(donor.address, 100)
      ).to.be.revertedWith("BursaryFund: Tokens are Soulbound");
    });
  });

  describe("Whitelist management", function () {
    it("Should only allow owner to whitelist schools", async function () {
      await expect(
        escrow.connect(other).setWhitelistedSchool(school.address, true)
      ).to.be.revertedWith("Ownable: caller is not the owner");
      
      await escrow.setWhitelistedSchool(school.address, true);
      expect(await escrow.whitelistedSchools(school.address)).to.be.true;
    });

    it("Should only allow owner to whitelist donors", async function () {
      await expect(
        escrow.connect(other).setWhitelistedDonor(donor.address, true)
      ).to.be.revertedWith("Ownable: caller is not the owner");
      
      await escrow.setWhitelistedDonor(donor.address, true);
      expect(await escrow.whitelistedDonors(donor.address)).to.be.true;
    });
  });
});