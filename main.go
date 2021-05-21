package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	filekeystore "github.com/ethersphere/bee/pkg/keystore/file"
	"github.com/pborman/uuid"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("exportSwarmKey <sourceDir> <password>")
		return
	}

	sourceDir := os.Args[1]
	password := os.Args[2]

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil || len(files) == 0 {
		fmt.Println("error reading source dir :, ", err.Error())
		return
	}
	fks := filekeystore.New(sourceDir)
	// sourceFile := filepath.Join(sourceDir, "swarm.key")
	privateKeyECDSA, _, err := fks.Key("swarm", password)
	if err != nil {
		fmt.Println("error reading key : ", err.Error())
		return
	}

	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}

	content, err := json.Marshal(key)
	if err != nil {
		fmt.Println("error marshalling key :", err.Error())
		return
	}
	fmt.Println(string(content))
	ioutil.WriteFile(filepath.Join(sourceDir, "private.json"), content, os.ModePerm)
}
