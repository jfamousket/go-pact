package pact

import "crypto/ed25519"

type KeyPair struct {
	Private ed25519.PrivateKey
	Public  ed25519.PublicKey
}
