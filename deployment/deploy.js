const etherlime = require('etherlime-lib');
const MerkleUtils = require('../build/MerkleUtils.json');
const ConvertUtils = require('../build/ConvertUtils.json');
const MerkleAirDrop = require('../build/MerkleAirDrop.json');
const ethers = require('ethers');
const utils = ethers.utils;


const deploy = async (network, secret) => {

	const deployer = new etherlime.EtherlimeGanacheDeployer();
	const MerkleUtilsLib = await deployer.deploy(MerkleUtils);
	const ConvertUtilsLib = await deployer.deploy(ConvertUtils);
	await deployer.deployAndVerify(MerkleAirDrop, { MerkleUtils: MerkleUtilsLib.contractAddress, ConvertUtils: ConvertUtilsLib.contractAddress });
	// const tx = await merkleContract.setRoot("0x1f2046f5ede7895de3666059b52edcc36e3fa4f8812bfd9ff34553f5aea45ec1")
	// await merkleContract.verboseWaitForTransaction(tx)
	// const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes("Gitcoin Livestream Rocks"), [
	// 	"0xac11234732f084af283c6f0abcd30bbab34de31fc1ae3040ae8b91cbe6a18794",
	// 	"0xf6c6901f1cd8f45d193642065a7b88f9d3549006be25adfd53cef07d8c6c434b"
	// ], 3)
	// console.log("This transaction is part of the Merkle tree: ", isPart);

};

module.exports = {
	deploy
};