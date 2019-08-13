package main

import (
	"./saver"
	"errors"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/LimeChain/merkletree/memory"
	"github.com/LimeChain/merkletree/postgres"
	merkleRestBaseAPI "github.com/LimeChain/merkletree/restapi/baseapi"
	merkleRestValidateAPI "github.com/LimeChain/merkletree/restapi/validateapi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/simonleung8/flags"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func createAndStartAPI(tree merkletree.ExternalMerkleTree, port int) {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.NoCache,
	)

	router.Route("/v1", func(r chi.Router) {
		treeRouter := chi.NewRouter()
		treeRouter = merkleRestBaseAPI.MerkleTreeStatus(treeRouter, tree)
		treeRouter = merkleRestBaseAPI.MerkleTreeInsert(treeRouter, tree)
		treeRouter = merkleRestBaseAPI.MerkleTreeHashes(treeRouter, tree)
		treeRouter = merkleRestValidateAPI.MerkleTreeValidate(treeRouter, tree)
		r.Mount("/api/merkletree", treeRouter)
	})

	fmt.Printf("Starting REST Api at port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func createSaver(tree merkletree.MerkleTree, nodeUrl, privateKeyHex, contractAddress string, periodSeconds int) {
	treeSaver, err := saver.NewSaver(
		nodeUrl,
		privateKeyHex,
		contractAddress,
		tree)
	if err != nil {
		panic(err)
	}

	go func() {
		len := 0
		timeout := time.Duration(periodSeconds) * time.Second
		for {
			savedRoot, err := treeSaver.FetchRoot()
			if err != nil {
				fmt.Println("ERR: Could not save the tree root")
				fmt.Println(err.Error())
				time.Sleep(timeout)
				continue
			}

			if savedRoot == tree.Root() {
				fmt.Printf("Same root (%v) found in the chain. Skipping this iteration\n", savedRoot)
				time.Sleep(timeout)
				continue
			}

			treeLen := tree.Length()
			if treeLen > len {
				fmt.Printf("Submitting new tree root to the chain (%v)\n", tree.Root())
				tx, err := treeSaver.TriggerSave()
				if err != nil {
					fmt.Println("ERR: Could not save the tree root")
					fmt.Println(err.Error())
				} else {
					fmt.Printf("Transaction (%v) mined\n", tx.TxHash.Hex())
					fmt.Printf("Gas used (%v)\n", tx.GasUsed)
					len = treeLen
				}
			} else {
				fmt.Println("No changes to submit. Skipping this iteration")
			}
			time.Sleep(timeout)
		}
	}()

	fmt.Printf("Started saver on %v seconds\n", periodSeconds)
	fmt.Printf("Node url %v\n", nodeUrl)
	fmt.Printf("Verifier contract address %v \n", contractAddress)

}

func setupFlags() flags.FlagContext {
	fc := flags.New()

	fc.NewBoolFlag("envirnonment-variables", "e", "Connection string for the postgres database")
	fc.NewStringFlag("database-connection", "db", "Connection string for the postgres database")
	fc.NewStringFlag("node-url", "n", "url to the ethereum node to connect")
	fc.NewStringFlag("secret", "s", "private key for the ethereum saver")
	fc.NewStringFlag("conntract-address", "c", "Address to the verifier contract")
	fc.NewIntFlagWithDefault("port", "ap", "port to run the API on", 8080)
	fc.NewIntFlagWithDefault("period", "p", "period to try and save the new root", 15)

	fc.Parse(os.Args...)

	return fc
}

func loadPostgreTree(connStr string) merkletree.FullMerkleTree {
	tree := postgres.LoadMerkleTree(memory.NewMerkleTree(), connStr)
	fmt.Printf("Merkle tree loaded. Length : %v, Root : %v\n", tree.Length(), tree.Root())
	return tree
}

func main() {
	fc := setupFlags()

	var connStr, nodeUrl, privateKeyHex, contractAddress string
	var period, port int

	if fc.Bool("e") {

	} else {
		if !fc.IsSet("db") {
			panic(errors.New("No db flag set"))
		}

		if !fc.IsSet("n") {
			panic(errors.New("No node-url flag set"))
		}

		if !fc.IsSet("s") {
			panic(errors.New("No secret flag set"))
		}

		if !fc.IsSet("c") {
			panic(errors.New("No conntract-address flag set"))
		}

		connStr = fc.String("db")
		nodeUrl = fc.String("n")
		privateKeyHex = fc.String("s")
		contractAddress = fc.String("c")
		period = fc.Int("p")
		port = fc.Int("ap")
	}

	tree := loadPostgreTree(connStr)

	createSaver(tree, nodeUrl, privateKeyHex, contractAddress, period)

	createAndStartAPI(tree, port)
}
