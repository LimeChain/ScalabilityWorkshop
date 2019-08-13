package saver

import (
	RootValidator "../contracts"
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/LimeChain/merkletree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
)

type EthereumRootSaver struct {
	client          *ethclient.Client
	privateKey      *ecdsa.PrivateKey
	contractAddress string
	tree            merkletree.MerkleTree
	mutex           sync.RWMutex
}

func (saver *EthereumRootSaver) GetContractAddress() string {
	return saver.contractAddress
}
func (saver *EthereumRootSaver) GetSaverWalletAddress() (string, error) {
	publicKey := saver.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("Could not derive public key from the saved private key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

func waitForTxReceipt(client *ethclient.Client, txHash common.Hash, seconds int) (*types.Receipt, error) {

	for i := 0; i < seconds; i++ {
		receipt, _ := client.TransactionReceipt(context.Background(), txHash)
		if receipt != nil {
			return receipt, nil
		}
		time.Sleep(1 * time.Second)
	}

	return nil, errors.New("Could not wait for tx in the given seconds period")

}

func (saver *EthereumRootSaver) TriggerSave() (*types.Receipt, error) {
	saver.mutex.Lock()
	defer saver.mutex.Unlock()

	address := common.HexToAddress(saver.contractAddress)
	contract, err := RootValidator.NewRootValidator(address, saver.client)
	if err != nil {
		return nil, err
	}

	tx, err := contract.SetRoot(bind.NewKeyedTransactor(saver.privateKey), common.HexToHash(saver.tree.Root()))
	if err != nil {
		return nil, err
	}

	receipt, err := waitForTxReceipt(saver.client, tx.Hash(), 180)

	if err != nil {
		return nil, err
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return nil, errors.New("The transaction was reverted")
	}

	return receipt, nil
}

func (saver *EthereumRootSaver) FetchRoot() (string, error) {
	address := common.HexToAddress(saver.contractAddress)
	contract, err := RootValidator.NewRootValidator(address, saver.client)
	if err != nil {
		return "", err
	}

	result, err := contract.LimeRoot(nil)
	if err != nil {
		return "", err
	}

	empty := [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if bytes.Equal(result[:], empty[:]) {
		return "", nil
	}

	hexResult := common.Bytes2Hex(result[:])

	return "0x" + hexResult, nil
}

func NewSaver(host, privateKeyHex, contractAddressHex string, tree merkletree.MerkleTree) (saver *EthereumRootSaver, err error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)

	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	_, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("The private key could not produce a public key")
	}

	saver = &EthereumRootSaver{}
	saver.client = client
	saver.privateKey = privateKey
	saver.contractAddress = contractAddressHex
	saver.tree = tree

	return saver, nil
}
