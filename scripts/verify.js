const etherlime = require('etherlime-lib');
const MerkleAirDrop = require('../build/MerkleAirDrop.json');
const ethers = require('ethers');
const utils = ethers.utils;

const verify = async () => {

	const contractAddress = "0x78d8dfC659aA493877230A247BE11fDf6294Fd2C";
	const originalData = "0x795eff09b1fe788dc7e6824aa5221ad893fd465a:5000000000000000000";
	const index = 0;
	const price = "5000000000000000000";
	const hashes = [
		"0x1e4c54758eee5ca89d9d3415033da6d17beb240fe2e237cd648a9170d237640b"
	];

	const provider = new ethers.providers.InfuraProvider('rinkeby', 'H4UAAWyThMPs2WB9LsHD')
	const wallet = new ethers.Wallet('d723d3cdf932464de15845c0719ca13ce15e64c83625d86ddbfc217bd2ac5f5a', provider)
	const merkleContract = await etherlime.ContractAt(MerkleAirDrop, contractAddress, wallet, provider)
	const isPart = await merkleContract.verifyDataInState(utils.toUtf8Bytes(originalData), hashes, index)
	const isPart2 = await merkleContract.testInState(price, hashes, index);
	console.log("This transaction is part of the Merkle tree: ", isPart);
	console.log("This transaction is part of the Merkle tree 2: ", isPart2);
	const tx = await merkleContract.claim(price, hashes, index)

	console.log("Claim transaction hash:", tx.hash);

};

verify()