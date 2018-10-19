package ethtest

import (
	"bufio"
	"bytes"
	"encoding/json"

	"github.com/ethereum/go-ethereum/rpc"
)

const (
	url             = "http://127.0.0.1:7545"
	evmIncreaseTime = "evm_increaseTime"
	evmSnapshot     = "evm_snapshot"
	evmMine         = "evm_mine"
	evmRevert       = "evm_revert"
)

// EvmIncreaseTime : Jump forward in time. Takes one parameter, which is the amount of time to increase in seconds.
// Returns the total time adjustment, in seconds.
func EvmIncreaseTime(seconds int64) (int64, error) {
	var result int64
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return 0, err
	}
	if err := rpcClient.Call(&result, evmIncreaseTime, seconds); err != nil {
		return 0, err
	}

	json.NewEncoder(bufio.NewWriter(&bytes.Buffer{})).Encode(result)

	return result, nil
}

// EvmMine : Force a block to be mined. Takes one optional parameter, which is the timestamp
// a block should setup as the mining time. Mines a block independent of whether or not mining
// is started or stopped.
func EvmMine(timestamp int64) (string, error) {
	var result string
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return "", err
	}
	if timestamp > 0 {
		if err := rpcClient.Call(&result, evmMine, timestamp); err != nil {
			return "", err
		}
	} else {
		if err := rpcClient.Call(&result, evmMine); err != nil {
			return "", err
		}
	}

	json.NewEncoder(bufio.NewWriter(&bytes.Buffer{})).Encode(result)

	return result, nil
}

// EvmRevert Revert the state of the blockchain to a previous snapshot.
// Takes a single parameter, which is the snapshot id to revert to.
// If no snapshot id is passed it will revert to the latest snapshot. Returns true.
func EvmRevert(snapshotId string) (bool, error) {
	var result bool
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return false, err
	}
	if err := rpcClient.Call(&result, evmRevert, snapshotId); err != nil {
		return false, err
	}

	json.NewEncoder(bufio.NewWriter(&bytes.Buffer{})).Encode(result)

	return result, nil
}

// EvmSnapshot: Snapshot the state of the blockchain at the current block. Takes no parameters. Returns the integer id of the snapshot created.
func EvmSnapshot() (string, error) {
	var result string
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return "", err
	}
	if err := rpcClient.Call(&result, evmSnapshot); err != nil {
		return "", err
	}

	json.NewEncoder(bufio.NewWriter(&bytes.Buffer{})).Encode(result)

	return result, nil
}
