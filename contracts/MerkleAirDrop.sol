pragma solidity ^0.5.2;

import "./RootValidator.sol";
import "./ConvertUtils.sol";

contract MerkleAirDrop is RootValidator {


	mapping(address => uint256) public balanceOf;
	mapping(address => bool) public hasClaimed;
	
	function claim(uint256 price, bytes32[] memory nodes, uint leafIndex) public {
		require(!hasClaimed[msg.sender]);
		bytes memory data = ConvertUtils.addressToBytesString(msg.sender);
		require(verifyDataInState(abi.encodePacked(data, ":", ConvertUtils.uintToBytesString(price)), nodes, leafIndex), "Data not contained");

		hasClaimed[msg.sender] = true;
		balanceOf[msg.sender] += price;
	}


	function testInState(uint256 price, bytes32[] memory nodes, uint leafIndex) view public returns(bool) {
		bytes memory data = ConvertUtils.addressToBytesString(address(msg.sender));
		return verifyDataInState(abi.encodePacked(data, ":", ConvertUtils.uintToBytesString(price)), nodes, leafIndex);
	}

}