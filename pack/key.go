package pack

import (
	_ "embed"
)

//go:embed public.key
var publicKey []byte

func PublicKey() []byte {
	return publicKey
}
