package zcncrypto

import (
	"encoding/json"
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"github.com/0chain/gosdk/core/encryption"
)

const cryptoVersion = "1.0"

// KeyPair private and publickey
type KeyPair struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// Wallet structure
type Wallet struct {
	ClientID    string    `json:"client_id"`
	ClientKey   string    `json:"client_key"`
	Keys        []KeyPair `json:"keys"`
	Mnemonic    string    `json:"mnemonics"`
	Version     string    `json:"version"`
	DateCreated string    `json:"date_created"`
}

//SignatureScheme - an encryption scheme for signing and verifying messages
type SignatureScheme interface {
	// Generate fresh keys
	GenerateKeys(numKeys int) (*Wallet, error)
	// Generate keys from mnemonic for recovery
	RecoverKeys(mnemonic string, numKeys int) (*Wallet, error)

	// Signing  - Set private key to sign
	SetPrivateKey(privateKey string) error
	Sign(hash string) (string, error)

	// Signature verification - Set public key to verify
	SetPublicKey(publicKey string) error
	Verify(signature string, msg string) (bool, error)

	// Combine signature for schemes BLS
	Add(signature, msg string) (string, error)
}

// NewSignatureScheme creates an instance for using signature functions
func NewSignatureScheme(sigScheme string) SignatureScheme {
	switch sigScheme {
	case "ed25519":
		return NewED255190chainScheme()
	case "bls0chain":
		return NewBLS0ChainScheme()
	default:
		panic(fmt.Sprintf("unknown signature scheme: %v", sigScheme))
	}
	return nil
}

// Marshal returns json string
func (w *Wallet) Marshal() (string, error) {
	ws, err := json.Marshal(w)
	if err != nil {
		return "", fmt.Errorf("Invalid Wallet")
	}
	return string(ws), nil
}

func IsMnemonicValid(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

func Sha3Sum256(data string) string {
	return encryption.Hash(data)
}