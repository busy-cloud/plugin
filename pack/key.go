package pack

import (
	"crypto/ed25519"
	"encoding/hex"
)

const PublicKey = "4f851cec1f93a757037fbb7771aead9a346df9cdd1cf623a8c00b691ac369ed5"
const PrivateKey = "fdf6dce03f6d7b9e8c6e51d99c2ab160aa9dc46ff839edaba794536aeb9335454f851cec1f93a757037fbb7771aead9a346df9cdd1cf623a8c00b691ac369ed5"

func PublicKeyBytes() []byte {
	var pubKey, _ = hex.DecodeString(PublicKey)
	return pubKey
}

func PrivateKeyBytes() []byte {
	var priKey, _ = hex.DecodeString(PrivateKey)
	return priKey
}

func GenerateKeyPair() (string, string, error) {
	pub, pri, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(pub), hex.EncodeToString(pri), err
}
