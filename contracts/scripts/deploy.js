const { ethers } = require("hardhat");

async function main() {
  // Deploy mock USDT token
  const usdtFactory = await ethers.getContractFactory("MockUSDT");
  const usdt = await usdtFactory.deploy();
  await usdt.deployed();
  console.log("USDT deployed to:", usdt.address);

  // Deploy BursaryEscrow
  const escrowFactory = await ethers.getContractFactory("BursaryEscrow");
  const escrow = await escrowFactory.deploy(usdt.address);
  await escrow.deployed();
  console.log("BursaryEscrow deployed to:", escrow.address);

  // Deploy VendorRegistry
  const registryFactory = await ethers.getContractFactory("VendorRegistry");
  const registry = await registryFactory.deploy();
  await registry.deployed();
  console.log("VendorRegistry deployed to:", registry.address);

  // Whitelist some test addresses
  const [deployer, donor1, donor2, school1, school2] = await ethers.getSigners();
  
  await escrow.setWhitelistedDonor(donor1.address, true);
  await escrow.setWhitelistedDonor(donor2.address, true);
  await escrow.setWhitelistedSchool(school1.address, true);
  await escrow.setWhitelistedSchool(school2.address, true);

  console.log("Test addresses whitelisted");
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });