package pact

type Scheme string

const (
	ED25519 Scheme = "ED25519"
	ETH     Scheme = "ETH"
)

type Meta struct {
	ChainId      string  `json:"chainId"`
	Sender       string  `json:"sender"`
	GasLimit     uint64  `json:"gasLimit"`
	GasPrice     float64 `json:"gasPrice"`
	Ttl          float64 `json:"ttl"`
	CreationTime uint64  `json:"creationTime"`
}

type Signer struct {
	Scheme Scheme `json:"scheme"`
	PubKey string `json:"pubKey"`
	Addr   string `json:"addr"`
	Caps   string `json:"caps"`
}

type Exec struct {
	Code string       `json:"code"`
	Data *interface{} `json:"data"`
}

func (Exec) PayloadType() {}

type Cont struct {
	PactId   string      `json:"pactId"`
	Rollback bool        `json:"rollback"`
	Step     uint64      `json:"step"`
	Data     interface{} `json:"data"`
	Proof    string      `json:"proof"`
}

func (Cont) PayloadType() {}

type PayloadType interface {
	PayloadType()
}

type Payload struct {
	Exec *Exec `json:"exec,omitempty"`
	Cont *Cont `json:"cont,omitempty"`
}

type CommandField struct {
	Nonce     string   `json:"nonce"`
	Meta      *Meta    `json:"meta"`
	Signers   []Signer `json:"signers"`
	Payload   Payload  `json:"payload"`
	NetworkId string   `json:"networkId,omitempty"`
}

type Sig struct {
	Sig string `json:"sig,omitempty"`
}

type Command struct {
	Cmd  string `json:"cmd"`
	Hash string `json:"hash"`
	Sigs []Sig  `json:"sigs"`
}

type SendCommand struct {
	Cmds []Command `json:"cmds,omitempty"`
}

type SPVCommand struct {
	RequestKey    string `json:"requestKey"`
	TargetChainId string `json:"targetChainId"`
}

type PactEvents struct {
	Name       string      `json:"name,omitempty"`
	Params     interface{} `json:"params,omitempty"`
	Module     string      `json:"module,omitempty"`
	ModuleHash string      `json:"moduleHash,omitempty"`
}

type Result struct {
	Status string      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
