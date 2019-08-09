const etherlime = require('etherlime-lib');
const ethers = require('ethers');
const DemoToken = require('../build/DemoToken.json');


const deploy = async (network, secret, etherscanApiKey) => {

	const deployer = new etherlime.EtherlimeGanacheDeployer();
	const result = await deployer.deploy(DemoToken, {}, ethers.utils.formatBytes32String("Placeholder"), ethers.utils.toUtf8Bytes("Qmcu7QgaAt28skUNFje1Nb6CyvWt888meddyBAipWKpBDT"));

};

module.exports = {
	deploy
};