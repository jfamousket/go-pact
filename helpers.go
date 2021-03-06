package pact

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/islishude/bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

type PrepareCommand interface {
	isPrepareCommand()
}

type PrepareExec struct {
	KeyPairs  []KeyPair
	Nonce     string
	EnvData   interface{}
	Meta      *Meta
	NetworkId string
	PactCode  string
}

func (c PrepareExec) isPrepareCommand() {}

type PrepareCont struct {
	KeyPairs  []KeyPair
	Nonce     string
	Proof     string
	Rollback  bool
	Step      uint64
	PactId    string
	EnvData   interface{}
	Meta      *Meta
	NetworkId string
}

func (c PrepareCont) isPrepareCommand() {}

func CreateBlake2Hash(data []byte) ([32]byte, string) {
	hash := blake2b.Sum256(data)
	return hash, base64.URLEncoding.EncodeToString(hash[:])
}

func EnforceType(value, valueType interface{}, msg string) {
	if ok := reflect.TypeOf(value) == reflect.TypeOf(valueType); !ok {
		panic(Error(fmt.Sprintf("%s must be a %t: %s", msg, valueType, value)))
	}
}

func EnforceNoError(err error) {
	if ok := err == nil; !ok {
		panic(Error(err.Error()))
	}
}

func EnforceValid(valid bool, err error) {
	if !valid {
		panic(Error(err.Error()))
	}
}

func MakeMeta(
	sender, chainId string,
	gasPrice float64,
	gasLimit uint64,
	creationTime uint64,
	ttl float64,
) *Meta {
	EnforceType(sender, "string", "sender")
	EnforceType(chainId, "string", "chainId")
	EnforceType(gasLimit, uint64(10), "gasLimit")
	EnforceType(creationTime, uint64(10), "creationTime")
	EnforceType(gasPrice, float64(10), "gasPrice")
	EnforceType(ttl, float64(10), "ttl")
	return &Meta{
		ChainId:      chainId,
		Sender:       sender,
		GasLimit:     gasLimit,
		GasPrice:     gasPrice,
		Ttl:          ttl,
		CreationTime: creationTime,
	}
}

func MakeSigner(keyPair KeyPair) Signer {
	return Signer{
		PubKey: hex.EncodeToString(keyPair.Public),
		Scheme: ED25519,
	}
}

func MarshalBody(value interface{}) (b *bytes.Buffer, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	body, err := json.Marshal(value)
	EnforceNoError(err)
	b = bytes.NewBuffer(body)
	return
}

func GenKeyPair(password string) (keyPair KeyPair, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(Error)
		}
	}()
	entropy, err := bip39.NewEntropy(128)
	EnforceNoError(err)

	mnemonic, err := bip39.NewMnemonic(entropy)
	EnforceNoError(err)

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, password)
	EnforceNoError(err)

	fmt.Printf("Seed Phrase: %s", mnemonic)

	privKey := bip32.NewRootXPrv(seed)
	keyPair = KeyPair{
		Private: ed25519.PrivateKey(privKey.String()),
		Public:  privKey.PublicKey(),
	}
	return
}

// PrepareExecCommand creates a command that can be sent to the /exec pact server route
func PrepareExecCommand(cmd PrepareExec) Command {
	EnforceType(cmd.Nonce, "string", "nonce")
	EnforceType(cmd.PactCode, "string", "pactCode")

	_ = getSigners(cmd)
	cmd.Meta = getMeta(cmd)
	cmd.Nonce = getCmdNonce(cmd)

	exec := Exec{
		Data: &cmd.EnvData,
		Code: cmd.PactCode,
	}

	cmdToSend, err := MarshalBody(CommandField{
		Signers: []Signer{},
		Meta:    cmd.Meta,
		Nonce:   cmd.Nonce,
		Payload: Payload{
			Exec: &exec,
			Cont: nil,
		},
		NetworkId: cmd.NetworkId,
	})
	EnforceNoError(err)

	return makeSingleCmd(cmdToSend.Bytes())
}

func PrepareContCmd(cmd PrepareCont) Command {
	EnforceType(cmd.Nonce, "string", "nonce")

	_ = getSigners(cmd)
	cmd.Nonce = getCmdNonce(cmd)
	cmd.Meta = getMeta(cmd)

	cont := Cont{
		Proof:    cmd.Proof,
		PactId:   cmd.PactId,
		Rollback: cmd.Rollback,
		Step:     cmd.Step,
		Data:     cmd.EnvData,
	}

	cmdToSend, err := MarshalBody(CommandField{
		Nonce:   cmd.Nonce,
		Meta:    cmd.Meta,
		Signers: []Signer{},
		Payload: Payload{
			Cont: &cont,
		},
		NetworkId: cmd.NetworkId,
	})

	EnforceNoError(err)

	return makeSingleCmd(cmdToSend.Bytes())
}

func getCmdNonce(cmd PrepareCommand) string {
	switch c := cmd.(type) {
	case PrepareCont:
		if c.Nonce == "" {
			return time.Now().Format(time.RFC3339)
		}
		return c.Nonce
	case PrepareExec:
		if c.Nonce == "" {
			return time.Now().Format(time.RFC3339)
		}
		return c.Nonce
	}
	return ""
}

func getSigners(cmd PrepareCommand) []Signer {
	var signers []Signer
	switch c := cmd.(type) {
	case PrepareCont:
		for _, kp := range c.KeyPairs {
			signers = append(signers, MakeSigner(kp))
		}
	case PrepareExec:
		for _, kp := range c.KeyPairs {
			signers = append(signers, MakeSigner(kp))
		}
	}
	return signers
}

func getMeta(cmd PrepareCommand) *Meta {
	switch c := cmd.(type) {
	case PrepareExec:
		if c.Meta == nil {
			return MakeMeta("", "0", 1e-9, 600, 0, 28800)
		}
		return c.Meta
	case PrepareCont:
		if c.Meta == nil {
			return MakeMeta("", "0", 1e-9, 600, 0, 28800)
		}
		return c.Meta
	}
	return nil
}

func makeSingleCmd(cmd []byte) Command {

	_, hash := CreateBlake2Hash(cmd)

	return Command{
		Cmd:  string(cmd),
		Hash: hash,
		Sigs: []Sig{},
	}
}
