const etherlime = require('etherlime-lib');
const MerkleUtils = require('../build/MerkleUtils.json');
const ConvertUtils = require('../build/ConvertUtils.json');
const MerkleAirDrop = require('../build/MerkleAirDrop.json');
const ethers = require('ethers');
const utils = ethers.utils;


const deploy = async (network, secret, etherscanApiKey) => {

	const deployer = new etherlime.InfuraPrivateKeyDeployer(secret, network, 'H4UAAWyThMPs2WB9LsHD', {
		etherscanApiKey
	})
	const MerkleUtilsLib = await deployer.deploy(MerkleUtils);
	const ConvertUtilsLib = await deployer.deploy(ConvertUtils);
	await deployer.deployAndVerify(MerkleAirDrop, { MerkleUtils: MerkleUtilsLib.contractAddress, ConvertUtils: ConvertUtilsLib.contractAddress });
};

module.exports = {
	deploy
};