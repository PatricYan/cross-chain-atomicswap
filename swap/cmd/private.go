/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/seeleteam/go-seele/cmd/util"
	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/common/hexutil"
	"github.com/seeleteam/go-seele/core/types"
	"github.com/seeleteam/go-seele/crypto"
	"github.com/seeleteam/go-seele/p2p"
	"github.com/seeleteam/go-seele/rpc2"
	"github.com/urfave/cli"
	"github.com/yanllearnn/cross-chain-atomicswap/contract"
)

func GetTxPoolContentAction(client *rpc.Client) (interface{}, error) {
	var result map[string][]map[string]interface{}
	err := client.Call(&result, "debug_getTxPoolContent")
	return result, err
}

func GetTxPoolTxCountAction(client *rpc.Client) (interface{}, error) {
	var result uint64
	err := client.Call(&result, "debug_getTxPoolTxCount")
	return result, err
}

func StartMinerAction(client *rpc.Client) (interface{}, error) {
	var result bool
	err := client.Call(&result, "miner_start", threadsValue)
	return result, err
}

func StopMinerAction(client *rpc.Client) (interface{}, error) {
	var result bool
	err := client.Call(&result, "miner_stop")
	return result, err
}

func SetMinerThreadsAction(client *rpc.Client) (interface{}, error) {
	var result bool
	err := client.Call(&result, "miner_setThreads", threadsValue)
	return result, err
}

func GetMinerThreadsAction(client *rpc.Client) (interface{}, error) {
	var result int
	err := client.Call(&result, "miner_getThreads")
	return result, err
}

func GetMinerStatusAction(client *rpc.Client) (interface{}, error) {
	var result string
	err := client.Call(&result, "miner_status")
	return result, err
}

func GetMinerHashrateAction(client *rpc.Client) (interface{}, error) {
	var result uint64
	err := client.Call(&result, "miner_hashrate")
	return result, err
}

func SetMinerCoinbaseAction(client *rpc.Client) (interface{}, error) {
	var result bool
	err := client.Call(&result, "miner_setCoinbase", coinbaseValue)
	return result, err
}

func GetMinerCoinbaseAction(client *rpc.Client) (interface{}, error) {
	var result common.Address
	err := client.Call(&result, "miner_getCoinbase")
	return result, err
}

func GetBlockTransactionCountAction(client *rpc.Client) (interface{}, error) {
	var result int
	var err error

	if hashValue != "" {
		err = client.Call(&result, "txpool_getBlockTransactionCountByHash", hashValue)
	} else {
		err = client.Call(&result, "txpool_getBlockTransactionCountByHeight", heightValue)
	}

	return result, err
}

func GetTransactionAction(client *rpc.Client) (interface{}, error) {
	var result map[string]interface{}
	var err error

	if hashValue != "" {
		err = client.Call(&result, "txpool_getTransactionByBlockHashAndIndex", hashValue, indexValue)
	} else {
		err = client.Call(&result, "txpool_getTransactionByBlockHeightAndIndex", heightValue, indexValue)
	}

	return result, err
}

func GetReceiptByTxHashAction(client *rpc.Client) (interface{}, error) {
	var result map[string]interface{}
	err := client.Call(&result, "txpool_getReceiptByTxHash", hashValue)
	return result, err
}

func GetTransactionByHashAction(client *rpc.Client) (interface{}, error) {
	return util.GetTransactionByHash(client, hashValue)
}

func GetPendingTransactionsAction(client *rpc.Client) (interface{}, error) {
	var result []map[string]interface{}
	err := client.Call(&result, "txpool_getPendingTransactions")
	return result, err
}

func GetPeerCountAction(client *rpc.Client) (interface{}, error) {
	var result int
	err := client.Call(&result, "network_getPeerCount")
	return result, err
}

// GetPeersInfo get peers information
func GetPeersInfo(client *rpc.Client) (interface{}, error) {
	var result []p2p.PeerInfo
	err := client.Call(&result, "network_getPeersInfo")
	return result, err
}

// GetNetworkVersion get current network version
func GetNetworkVersion(client *rpc.Client) (interface{}, error) {
	var result uint64
	err := client.Call(&result, "network_getNetworkVersion")
	return result, err
}

// GetProtocolVersion get seele protocol version
func GetProtocolVersion(client *rpc.Client) (interface{}, error) {
	var result uint
	err := client.Call(&result, "network_getProtocolVersion")
	return result, err
}

// GetDumpHeap dump heap for profiling
func GetDumpHeap(client *rpc.Client) (interface{}, error) {
	var result string
	err := client.Call(&result, "debug_dumpHeap", dumpFileValue, gcBeforeDump)
	return result, err
}

// Swap is swap of seele and eth
func Swap(c *cli.Context) error {
	aliceCoinbase := "0xe5cf366f4c96ffecc662e5e0c4ead897d7c0d3d1"
	alicePrivateKey := "0x501440a7f9c5117ade294573d0609054f9fe515a91d2f282856d5df9dc619899"
	bobCoinbase := "0x20667d2cd25e332daaad4c71e6a7fc8db4738ac1"
	bobPrivateKey := "0x35d4b0ee697aa40b94e95648f3051e30b959d878a1403e543baa614d09b52b6d"
	client, err := rpc.DialTCP(context.Background(), addressValue)
	if err != nil {
		return err
	}

	bytecode, err := contract.GetContractByteCode()
	if err != nil {
		return err
	}

	// value set
	amountValue = "100000"
	feeValue = "100000"
	paloadValue = bytecode
	
	txd, err := checkParameterWithoutPassW(client, fromValue, "", amountValue, feeValue, 0)
	if err != nil {
		return err
	}

	// alice address privatekey for test no password
	privateKey, err := crypto.LoadECDSAFromString(alicePrivateKey)
	if err != nil {
		return err
	}

	tx, err := util.GenerateTx(privateKey, txd.To, txd.Amount, txd.Fee, txd.AccountNonce, txd.Payload)
	if err != nil {
		return err
	}

	var result bool
	err = client.Call(&result, "seele_addTx", tx)
	if err != nil {
		return err
	}

	contractAddr := crypto.CreateAddress(tx.Data.From, tx.Data.AccountNonce)
	fmt.Println("contract address:", contractAddr)

	var result1 map[string]interface{}
waitBlock:
	for {
		err := client.Call(&result1, "txpool_getReceiptByTxHash", tx.Hash.ToHex())
		if err != nil {
			time.Sleep(15)
			continue
		}
		break waitBlock
	}

	seelecontract, err := contract.NewSeeleContract(contractAddr)
	if err != nil {
		return err
	}

	var result1 map[string]interface{}
	err = client.Call(&result, "txpool_getTransactionByHash", tx.Hash.ToHex())
	if err != nil {
		return err
	}
	fmt.Println("result:", result1)

	// use newContract function in the deployed contract
	// secretlock := make([]byte, 32)
	// rand.Read(secretlock[:])
	// secrethash := sha256.Sum256(secretlock)
	timelock := time.Now().Unix() + 48*60*3600
	// fmt.Println("secrethash:", hexutil.BytesToHex(secrethash[:]))
	// secrethash hex: 0x129b2734e389b4b8f9b29733cb11cc6161b7f8332eee25c37b08af195f90738a
	// secretlock hex: 0x1913c75c159b4bd7f67d8f0319c30b313edfebe43fb9bfc7789b39a7fc5e969b
	// secrethash byte: [18 155 39 52 227 137 180 184 249 178 151 51 203 17 204 97 97 183 248 51 46 238 37 195 123 8 175 25 95 144 115 138]
	// secretlock byte: [25 19 199 92 21 155 75 215 246 125 143 3 25 195 11 49 62 223 235 228 63 185 191 199 120 155 57 167 252 94 150 155]
	secretHashHex := "0x129b2734e389b4b8f9b29733cb11cc6161b7f8332eee25c37b08af195f90738a"
	bytecode, err := seelecontract.NewContract(hexutil.HexToBytes(bobCoinbase), hexutil.HexToBytes(secretHashHex), timelock)
	if err != nil {
		return err
	}

	// audit

	// deploy contract in second seele chain using the bytecode

	return err
}

func sendtx(privateKey *ecdsa.PrivateKey, client *rpc.Client, fromHex, toHex, amountStr, feeStr string, nonceNum uint64, payload []byte) (*types.Transaction, error) {
	txd, err := checkParameterWithoutPassW(client, fromHex, toHex, amountStr, feeStr, nonceNum, payload)
	if err != nil {
		return nil, err
	}

	tx, err := util.GenerateTx(privateKey, txd.To, txd.Amount, txd.Fee, txd.AccountNonce, txd.Payload)
	if err != nil {
		return nil, err
	}

	var result bool
	err = client.Call(&result, "seele_addTx", tx)
	return tx, err
}
