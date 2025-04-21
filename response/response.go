package response

import (
	"encoding/binary"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Transaction struct {
	ChainId  int64          `json:"chainid"`
	EndPoint string         `json:"endpoint"`
	Nonce    uint64         `json:"nonce"`
	GasPrice *big.Int       `json:"gasprice"`
	Gas      uint64         `json:"gas"`
	To       common.Address `json:"to"`
	Value    *big.Int       `json:"value"`
	Data     []byte         `json:"data"`
}

func (t *Transaction) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Transaction) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

type SignTx struct {
	EndPoint string `json:"endpoint"`
	Tx       []byte `json:"tx"`
}

func (s *SignTx) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SignTx) Unmarshal(data []byte) error {
	return json.Unmarshal(data, s)
}

type CheckResponse struct {
	PayAddr  common.Address
	Seller   common.Address
	SizeByte uint64
	Nonce    *big.Int
}

func (c *CheckResponse) Hash() []byte {
	tmp8 := make([]byte, 8)
	binary.BigEndian.PutUint64(tmp8, c.SizeByte)

	m := common.LeftPadBytes(c.Nonce.Bytes(), 32)

	hash := crypto.Keccak256(c.PayAddr.Bytes(), m, tmp8, c.Seller.Bytes())
	return hash
}
