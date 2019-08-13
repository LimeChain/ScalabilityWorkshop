pragma solidity ^0.5.2;

import "./MerkleUtils.sol";
import "@openzeppelin/contracts/ownership/Ownable.sol";

contract RootValidator is Ownable {

	bytes32 public limeRoot;

	function setRoot(bytes32 merkleRoot) public onlyOwner {
		limeRoot = merkleRoot;
	}

	function verifyDataInState(bytes memory data, bytes32[] memory nodes, uint leafIndex) view public returns(bool) {
		return MerkleUtils.containedInTree(limeRoot, data, nodes, leafIndex);
	}

    
}