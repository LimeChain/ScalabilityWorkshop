package main

type RootSaver interface {
	FetchRoot() string
	GetContractAddress() string
	GetSaverWalletAddress() string
	TriggerSave() (string, error)
}
