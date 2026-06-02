// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

contract VendorRegistry is Ownable {
    mapping(address => bool) public whitelistedSchools;
    mapping(address => bool) public whitelistedDonors;
    
    event SchoolWhitelisted(address indexed school, bool status);
    event DonorWhitelisted(address indexed donor, bool status);
    
    function addSchool(address school) external onlyOwner {
        whitelistedSchools[school] = true;
        emit SchoolWhitelisted(school, true);
    }
    
    function removeSchool(address school) external onlyOwner {
        whitelistedSchools[school] = false;
        emit SchoolWhitelisted(school, false);
    }
    
    function addDonor(address donor) external onlyOwner {
        whitelistedDonors[donor] = true;
        emit DonorWhitelisted(donor, true);
    }
    
    function removeDonor(address donor) external onlyOwner {
        whitelistedDonors[donor] = false;
        emit DonorWhitelisted(donor, false);
    }
    
    function isSchoolWhitelisted(address school) external view returns (bool) {
        return whitelistedSchools[school];
    }
    
    function isDonorWhitelisted(address donor) external view returns (bool) {
        return whitelistedDonors[donor];
    }
}