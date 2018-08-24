package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/common/hexutil"
)

// SeeleContract is a type of seelecontract to parse the function
type SeeleContract struct {
	Address common.Address // Deployment address of the contract on the seele blockchain
	Abi     abi.ABI        // Reflect based ABI to access the correct Ethereum methods
}

// NewSeeleContract Create a seelecontract
func NewSeeleContract(address common.Address) (*SeeleContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashedTimelockABI))
	if err != nil {
		return nil, err
	}

	return nil, &SeeleContract{
		Address: address,
		Abi:     parsed,
	}
}

// GetContractByteCode get the contract bytecode
func GetContractByteCode() ([]byte, error) {
	bytecode, err := hexutil.HexToBytes(HashedTimelockBin)
	if err != nil {
		return nil, err
	}

	return bytecode, err
}

// GetContract return the GetContract function bytecode
func (seele *SeeleContract) GetContract(_contractId [32]byte) ([]byte, error) {
	return getFuncByteCode("getContract", _contractId)
}

// NewContract return the NewContract function bytecode
func (seele *SeeleContract) NewContract(_receiver common.Address, _hashlock [32]byte, _timelock int64) ([]byte, error) {
	return getFuncByteCode("newContract", _receiver, _hashlock, _timelock)
}

// Refund return the Refund function bytecode
func (seele *SeeleContract) Refund(_contractId [32]byte) ([]byte, error) {
	return getFuncByteCode("refund", _contractId)
}

// Withdraw return the Withdraw function bytecode
func (seele *SeeleContract) Withdraw(_contractId [32]byte, _preimage [32]byte) ([]byte, error) {
	return getFuncByteCode("withdraw", _contractId, _preimage)
}

func (seele *SeeleContract) getFuncByteCode(method string, params ...interface{}) ([]byte, error) {
	return seele.Abi.Pack(method, params...)
}
