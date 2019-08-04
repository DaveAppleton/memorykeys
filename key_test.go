package memorykeys

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestGenKey(t *testing.T) {
	key, err := GetPrivateKey("banker")
	if err != nil {
		t.Error("1", err)
	}
	key2, err := GetPrivateKey("banker")
	if err != nil {
		t.Error("2", err)
	}
	pub := key.PublicKey
	pub2 := key2.PublicKey
	if pub != pub2 {
		t.Error(crypto.PubkeyToAddress(pub), " != ", crypto.PubkeyToAddress(pub2))
	}
}

func TestImport(t *testing.T) {
	privateKey := "d31a46c5322e8e8a7e11f51cf9c4073fea42d33b431b5e7e76a82518fc178ea8"
	key, err := ImportPrivateKey("imported", privateKey)
	if err != nil {
		t.Error(err)
	}
	key2, err := GetPrivateKey("imported")
	if err != nil {
		t.Error(err)
	}
	pub := key.PublicKey
	pub2 := key2.PublicKey
	if pub != pub2 {
		t.Error(crypto.PubkeyToAddress(pub), " != ", crypto.PubkeyToAddress(pub2))
	}
	address := "0xbade08f05BCaecBe15924621E81Ead37f7895F0c"
	if crypto.PubkeyToAddress(pub).Hex() != address {
		t.Error("Incorrect address generated : ", crypto.PubkeyToAddress(pub).Hex(), " expecting ", address)
	}
}
