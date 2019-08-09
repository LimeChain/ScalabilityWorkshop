const etherlime = require('etherlime-lib');
const DemoToken = require('../build/DemoToken.json');


describe('Example', () => {
    let aliceAccount = accounts[3];
    let deployer;
    let demoTokenInstance;
    const name = ethers.utils.formatBytes32String("Placeholder");
    const image = ethers.utils.hexlify(ethers.utils.toUtf8Bytes("Qmcu7QgaAt28skUNFje1Nb6CyvWt888meddyBAipWKpBDT"));

    beforeEach(async () => {
        deployer = new etherlime.EtherlimeGanacheDeployer(aliceAccount.secretKey);
        demoTokenInstance = await deployer.deploy(DemoToken, {}, ethers.utils.formatBytes32String("Placeholder"), ethers.utils.toUtf8Bytes("Qmcu7QgaAt28skUNFje1Nb6CyvWt888meddyBAipWKpBDT"));
    });

    it('should have set the owner and data of the first token successfully', async () => {
        const placeholderItem = await demoTokenInstance.items(0);
        assert.strictEqual(placeholderItem.name, name, "Name was not set correctly");
        assert.strictEqual(placeholderItem.image, image, "Image was not set correctly");

        const owner = await demoTokenInstance.ownerOf(0);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not set to alice");
    });


});