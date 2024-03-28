package common

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Method string

const (
	AskStateBatch  Method = "askStateBatch"
	SignStateBatch Method = "signStateBatch"
	AskSlash       Method = "askSlash"
	SignSlash      Method = "signSlash"
	SignRollBack   Method = "signRollBack"
	AskRollBack    Method = "askRollBack"

	SlashTypeLiveness byte = 0
	SlashTypeCulprit  byte = 1

	CulpritErrorCode = 100
)

func (m Method) String() string {
	return string(m)
}

type SignStateRequest struct {
	Type                uint64     `json:"type"`
	StartBlock          *big.Int   `json:"start_block"`
	OffsetStartsAtIndex *big.Int   `json:"offset_starts_at_index"`
	Challenge           string     `json:"challenge"`
	StateRoots          [][32]byte `json:"state_roots"`
	ElectionId          uint64     `json:"election_id"`
}

func (ssr SignStateRequest) String() string {
	var srs string
	for _, sr := range ssr.StateRoots {
		srs = srs + hex.EncodeToString(sr[:]) + " "
	}
	return fmt.Sprintf("start_block: %v, offset_starts_at_index: %v, election_id: %d, state_roots: %s", ssr.StartBlock, ssr.OffsetStartsAtIndex, ssr.ElectionId, srs)
}

type SlashRequest struct {
	Address    common.Address `json:"address"`
	BatchIndex uint64         `json:"batch_index"`
	SignType   byte           `json:"sign_type"`
}

type RollBackRequest struct {
	StartBlock *big.Int `json:"start_block"`
}

type AskResponse struct {
	Result bool `json:"result"`
}

type NodeSignRequest struct {
	ClusterPublicKey string      `json:"cluster_public_key"`
	Timestamp        int64       `json:"timestamp"`
	Nodes            []string    `json:"nodes"`
	RequestBody      interface{} `json:"request_body"`
}

type SignResponse struct {
	Signature []byte `json:"signature"`
}

type KeygenRequest struct {
	Nodes      []string `json:"nodes"`
	ElectionId uint64   `json:"election_id"`
	Threshold  int      `json:"threshold"`
	Timestamp  int64    `json:"timestamp"`
}

type KeygenResponse struct {
	ClusterPublicKey string `json:"cluster_public_key"`
}

type SignatureData struct {
	// Ethereum-style recovery byte; only the first byte is relevant
	SignatureRecovery []byte `json:"signature_recovery,omitempty"`
	// Signature components R, S
	R []byte `json:"r,omitempty"`
	S []byte `json:"s,omitempty"`
	// M represents the original message digest that was signed M
	M []byte `json:"m,omitempty"`
}

type BatchSubmitterResponse struct {
	Signature []byte `json:"signature"`
	RollBack  bool   `json:"roll_back"`
}
