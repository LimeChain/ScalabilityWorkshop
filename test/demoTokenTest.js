const etherlime = require('etherlime-lib');
const ethers = require('ethers');
const DemoToken = require('../build/DemoToken.json');


describe('Example', () => {
    let aliceAccount = accounts[3];
    let bobsAccount = accounts[4];
    let carolAccount = accounts[5];
    let deployer;
    let demoTokenInstance;
    const name = ethers.utils.formatBytes32String("Placeholder");
    const image = ethers.utils.hexlify(ethers.utils.toUtf8Bytes("Qmcu7QgaAt28skUNFje1Nb6CyvWt888meddyBAipWKpBDT"));
    const ETH = ethers.utils.parseEther("1");

    const signMessage = (wallet, id, receiver, price, nonce) => {

        const hashMsg = ethers.utils.solidityKeccak256(['uint256', 'address', 'uint256', 'uint256'], [id, receiver, price, nonce]);
        var hashData = ethers.utils.arrayify(hashMsg);
        const signature = wallet.signMessage(hashData);
        return signature;
    }

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


    it('should be able to transfer with signed message', async () => {
        const id = 0;
        const nonce = await demoTokenInstance.nonceOf(id);
        const signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH });

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, bobsAccount.signer.address, "Owner was not set to bob");
    });

    it('should be able to transfer with signed message once again', async () => {
        const id = 0;
        let nonce = await demoTokenInstance.nonceOf(id);
        let signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH });

        nonce = await demoTokenInstance.nonceOf(id);
        signature = signMessage(bobsAccount.signer, id, carolAccount.signer.address, ETH, nonce);
        await demoTokenInstance.safeTransfer(id, carolAccount.signer.address, ETH, signature, { value: ETH });

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, carolAccount.signer.address, "Owner was not set to carol");

    });

    it('should not be able to transfer without correct signed message', async () => {
        const id = 0;
        const nonce = await demoTokenInstance.nonceOf(id);
        const signature = signMessage(carolAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await assert.revert(demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH }));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not still alice");
    });

    it('should not be able to transfer without correct nonce', async () => {
        const id = 0;
        const signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, 14);
        await assert.revert(demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH }));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not still alice");
    });

    it('should not be able to transfer without correct receiver', async () => {
        const id = 0;
        const nonce = await demoTokenInstance.nonceOf(id);
        const signature = signMessage(aliceAccount.signer, id, carolAccount.signer.address, ETH, nonce);
        await assert.revert(demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH }));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not still alice");
    });

    it('should not be able to transfer without correct price', async () => {
        const id = 0;
        const nonce = await demoTokenInstance.nonceOf(id);
        const signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await assert.revert(demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, '100', signature, { value: ETH }));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not still alice");
    });

    it('should not be able to transfer without sending price', async () => {
        const id = 0;
        const nonce = await demoTokenInstance.nonceOf(id);
        const signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await assert.revert(demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, aliceAccount.signer.address, "Owner was not still alice");
    });

    it('should not be able to transfer from the previous owner', async () => {
        const id = 0;
        let nonce = await demoTokenInstance.nonceOf(id);
        let signature = signMessage(aliceAccount.signer, id, bobsAccount.signer.address, ETH, nonce);
        await demoTokenInstance.safeTransfer(id, bobsAccount.signer.address, ETH, signature, { value: ETH });

        nonce = await demoTokenInstance.nonceOf(id);
        signature = signMessage(aliceAccount.signer, id, carolAccount.signer.address, ETH, nonce);
        await assert.revert(demoTokenInstance.safeTransfer(id, carolAccount.signer.address, ETH, signature, { value: ETH }));

        const owner = await demoTokenInstance.ownerOf(id);
        assert.strictEqual(owner, bobsAccount.signer.address, "Owner was not set to bob");
    });


});