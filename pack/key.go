package pack

import (
	"crypto/ed25519"
	"encoding/hex"
)

const PublicKey = "4f851cec1f93a757037fbb7771aead9a346df9cdd1cf623a8c00b691ac369ed5"

//const PrivateKey = ""

func PublicKeyBytes() []byte {
	var pubKey, _ = hex.DecodeString(PublicKey)
	return pubKey
}

func GenerateKeyPair() (string, string, error) {
	pub, pri, err := ed25519.GenerateKey(nil)
	if err != nil {
		return "", "", err
	}
	return hex.EncodeToString(pub), hex.EncodeToString(pri), err
}
