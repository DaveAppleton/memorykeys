package memorykeys

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var keyMap = make(map[string]*ecdsa.PrivateKey)

// GetPrivateKey get the private key mapped to the name. Generate the key if it does not exist
func GetPrivateKey(keyName string) (*ecdsa.PrivateKey, error) {
	var err error
	if keyMap[keyName] == nil {
		keyMap[keyName], err = crypto.GenerateKey()
	}
	return keyMap[keyName], err
}

// GetAddress get the ethereum address tied to the key
func GetAddress(keyName string) (*common.Address, error) {
	key, err := GetPrivateKey(keyName)
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(key.PublicKey)
	return &address, nil
}

// GetTransactor get a transactor for use with
func GetTransactor(keyName string) (*bind.TransactOpts, error) {
	key, err := GetPrivateKey(keyName)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactor(key), nil
}

// ImportPrivateKey use a specific private key (hex string) for the account
// use with caution - do not reveal keys in repos!!!
func ImportPrivateKey(keyName string, hexKey string) (*ecdsa.PrivateKey, error) {
	key, err := crypto.HexToECDSA(strings.TrimPrefix(hexKey, "0x"))
	if err == nil {
		keyMap[keyName] = key
	}
	return key, err
}
