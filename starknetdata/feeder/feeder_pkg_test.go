package feeder

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/NethermindEth/juno/clients/feeder"
	"github.com/NethermindEth/juno/core"
	"github.com/NethermindEth/juno/core/felt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Todo: refactor all of the test such that the test functions call the functions defined in
// 	StarknetData interface instead of calling adapt functions.

func TestAdaptBlock(t *testing.T) {
	block11817Json, err := os.ReadFile("../../clients/feeder/testdata/mainnet/block/11817.json")
	assert.NoError(t, err)
	block147Json, err := os.ReadFile("../../clients/feeder/testdata/mainnet/block/147.json")
	assert.NoError(t, err)

	t.Run("mainnet block number 11817", func(t *testing.T) {
		var response *feeder.Block
		err := json.Unmarshal(block11817Json, &response)
		require.NoError(t, err)

		expectedEventCount := uint64(0)
		for _, r := range response.Receipts {
			expectedEventCount = expectedEventCount + uint64(len(r.Events))
		}

		block, err := AdaptBlock(response)
		require.NoError(t, err)

		assert.True(t, block.Hash.Equal(response.Hash))
		assert.True(t, block.ParentHash.Equal(response.ParentHash))
		assert.Equal(t, response.Number, block.Number)
		assert.True(t, block.GlobalStateRoot.Equal(response.StateRoot))
		assert.Equal(t, response.Timestamp, block.Timestamp)
		assert.Equal(t, len(response.Transactions), len(block.Transactions))
		assert.Equal(t, uint64(len(response.Transactions)), block.TransactionCount)
		assert.Equal(t, len(response.Receipts), len(block.Receipts))
		assert.Equal(t, expectedEventCount, block.EventCount)
		assert.Equal(t, "0.10.1", block.ProtocolVersion)
		assert.Nil(t, block.ExtraData)
	})
	t.Run("mainnet block number 147", func(t *testing.T) {
		var response *feeder.Block
		err := json.Unmarshal(block147Json, &response)
		require.NoError(t, err)

		expectedEventCount := uint64(0)
		for _, r := range response.Receipts {
			expectedEventCount = expectedEventCount + uint64(len(r.Events))
		}

		block, err := AdaptBlock(response)
		require.NoError(t, err)

		assert.True(t, block.Hash.Equal(response.Hash))
		assert.True(t, block.ParentHash.Equal(response.ParentHash))
		assert.Equal(t, response.Number, block.Number)
		assert.True(t, block.GlobalStateRoot.Equal(response.StateRoot))
		assert.Equal(t, response.Timestamp, block.Timestamp)
		assert.Equal(t, len(response.Transactions), len(block.Transactions))
		assert.Equal(t, uint64(len(response.Transactions)), block.TransactionCount)
		assert.Equal(t, len(response.Receipts), len(block.Receipts))
		assert.Equal(t, expectedEventCount, block.EventCount)
		assert.Equal(t, "", block.ProtocolVersion)
		assert.Nil(t, block.ExtraData)
	})
	t.Run("error with unknown transaction", func(t *testing.T) {
		var response *feeder.Block
		err := json.Unmarshal(block147Json, &response)
		require.NoError(t, err)

		response.Transactions[0].Type = "test"

		block, err := AdaptBlock(response)
		assert.Nil(t, block)
		assert.EqualError(t, err, "unknown transaction")
	})
}

func hexToFelt(t *testing.T, hex string) *felt.Felt {
	f, err := new(felt.Felt).SetString(hex)
	require.NoError(t, err)
	return f
}

func TestAdaptStateUpdate(t *testing.T) {
	jsonData := []byte(`{
  "block_hash": "0x3",
  "new_root": "0x1",
  "old_root": "0x2",
  "state_diff": {
    "storage_diffs": {
      "0x20cfa74ee3564b4cd5435cdace0f9c4d43b939620e4a0bb5076105df0a626c6": [
        {
          "key": "0x5",
          "value": "0x22b"
        },
        {
          "key": "0x2",
          "value": "0x1"
        }
      ],
      "0x3": [
        {
          "key": "0x1",
          "value": "0x5"
        },
        {
          "key": "0x7",
          "value": "0x13"
        }
      ]
    },
    "nonces": { 
		"0x37" : "0x44",
		"0x44" : "0x37"
	},
    "deployed_contracts": [
      {
        "address": "0x1",
        "class_hash": "0x2"
      },
      {
        "address": "0x3",
        "class_hash": "0x4"
      }
	],
    "declared_contracts": [
		"0x37", "0x44"
	]
  }
}`)

	var feederStateUpdate feeder.StateUpdate
	err := json.Unmarshal(jsonData, &feederStateUpdate)
	assert.Equal(t, nil, err, "Unexpected error")

	coreStateUpdate, err := AdaptStateUpdate(&feederStateUpdate)
	if assert.NoError(t, err) {
		assert.Equal(t, true, feederStateUpdate.NewRoot.Equal(coreStateUpdate.NewRoot))
		assert.Equal(t, true, feederStateUpdate.OldRoot.Equal(coreStateUpdate.OldRoot))
		assert.Equal(t, true, feederStateUpdate.BlockHash.Equal(coreStateUpdate.BlockHash))

		assert.Equal(t, 2, len(feederStateUpdate.StateDiff.DeclaredContracts))
		for idx := range feederStateUpdate.StateDiff.DeclaredContracts {
			gw := feederStateUpdate.StateDiff.DeclaredContracts[idx]
			core := coreStateUpdate.StateDiff.DeclaredClasses[idx]
			assert.Equal(t, true, gw.Equal(core))
		}

		for keyStr, gw := range feederStateUpdate.StateDiff.Nonces {
			key, _ := new(felt.Felt).SetString(keyStr)
			core := coreStateUpdate.StateDiff.Nonces[*key]
			assert.Equal(t, true, gw.Equal(core))
		}

		assert.Equal(t, 2, len(feederStateUpdate.StateDiff.DeployedContracts))
		for idx := range feederStateUpdate.StateDiff.DeployedContracts {
			gw := feederStateUpdate.StateDiff.DeployedContracts[idx]
			core := coreStateUpdate.StateDiff.DeployedContracts[idx]
			assert.Equal(t, true, gw.ClassHash.Equal(core.ClassHash))
			assert.Equal(t, true, gw.Address.Equal(core.Address))
		}

		assert.Equal(t, 2, len(feederStateUpdate.StateDiff.StorageDiffs))
		for keyStr, diffs := range feederStateUpdate.StateDiff.StorageDiffs {
			key, _ := new(felt.Felt).SetString(keyStr)
			coreDiffs := coreStateUpdate.StateDiff.StorageDiffs[*key]
			assert.Equal(t, true, len(diffs) > 0)
			assert.Equal(t, len(diffs), len(coreDiffs))
			for idx := range diffs {
				assert.Equal(t, true, diffs[idx].Key.Equal(coreDiffs[idx].Key))
				assert.Equal(t, true, diffs[idx].Value.Equal(coreDiffs[idx].Value))
			}
		}
	}
}

func TestAdaptClass(t *testing.T) {
	classJson, err := os.ReadFile("../../clients/feeder/testdata/goerli/class/0x1924aa4b0bedfd884ea749c7231bafd91650725d44c91664467ffce9bf478d0.json")
	assert.NoError(t, err)

	response := new(feeder.ClassDefinition)
	err = json.Unmarshal(classJson, response)
	assert.NoError(t, err)

	class, err := adaptClass(response)
	assert.NoError(t, err)

	for i, v := range response.EntryPoints.External {
		assert.Equal(t, v.Selector, class.Externals[i].Selector)
		assert.Equal(t, v.Offset, class.Externals[i].Offset)
	}
	assert.Equal(t, len(response.EntryPoints.External), len(class.Externals))

	for i, v := range response.EntryPoints.L1Handler {
		assert.Equal(t, v.Selector, class.L1Handlers[i].Selector)
		assert.Equal(t, v.Offset, class.L1Handlers[i].Offset)
	}
	assert.Equal(t, len(response.EntryPoints.L1Handler), len(class.L1Handlers))

	for i, v := range response.EntryPoints.Constructor {
		assert.Equal(t, v.Selector, class.Constructors[i].Selector)
		assert.Equal(t, v.Offset, class.Constructors[i].Offset)
	}
	assert.Equal(t, len(response.EntryPoints.Constructor), len(class.Constructors))

	for i, v := range response.Program.Builtins {
		assert.Equal(t, new(felt.Felt).SetBytes([]byte(v)), class.Builtins[i])
	}
	assert.Equal(t, len(response.Program.Builtins), len(class.Builtins))

	for i, v := range response.Program.Data {
		expected, err := new(felt.Felt).SetString(v)
		assert.NoError(t, err)
		assert.Equal(t, expected, class.Bytecode[i])
	}
	assert.Equal(t, len(response.Program.Data), len(class.Bytecode))

	programHash, err := feeder.ProgramHash(response)
	assert.NoError(t, err)
	assert.Equal(t, programHash, class.ProgramHash)
}

func TestAdaptTransaction(t *testing.T) {
	t.Run("invoke transaction", func(t *testing.T) {
		invokeJson, err := os.ReadFile("../../clients/feeder/testdata/goerli/transaction/0x7e3a229febf47c6edfd96582d9476dd91a58a5ba3df4553ae448a14a2f132d9.json")
		assert.NoError(t, err)

		response := new(feeder.TransactionStatus)
		err = json.Unmarshal(invokeJson, response)
		require.NoError(t, err)

		transaction := response.Transaction
		txn, err := adaptTransaction(transaction)
		invokeTx, ok := txn.(*core.InvokeTransaction)
		require.True(t, ok)
		require.NoError(t, err)

		assert.Equal(t, transaction.Hash, invokeTx.Hash())
		assert.Equal(t, transaction.ContractAddress, invokeTx.ContractAddress)
		assert.Equal(t, transaction.EntryPointSelector, invokeTx.EntryPointSelector)
		assert.Equal(t, transaction.Nonce, invokeTx.Nonce)
		assert.Equal(t, transaction.CallData, invokeTx.CallData)
		assert.Equal(t, transaction.Signature, invokeTx.Signature())
		assert.Equal(t, transaction.MaxFee, invokeTx.MaxFee)
		assert.Equal(t, transaction.Version, invokeTx.Version)
	})
	t.Run("deploy transaction", func(t *testing.T) {
		deployJson, err := os.ReadFile("../../clients/feeder/testdata/goerli/transaction/0x15b51c2f4880b1e7492d30ada7254fc59c09adde636f37eb08cdadbd9dabebb.json")
		assert.NoError(t, err)

		response := new(feeder.TransactionStatus)
		err = json.Unmarshal(deployJson, response)
		require.NoError(t, err)

		transaction := response.Transaction
		txn, err := adaptTransaction(transaction)
		deployTx, ok := txn.(*core.DeployTransaction)
		require.True(t, ok)
		require.NoError(t, err)

		assert.Equal(t, transaction.Hash, deployTx.Hash())
		assert.Equal(t, transaction.ContractAddressSalt, deployTx.ContractAddressSalt)
		assert.Equal(t, transaction.ContractAddress, deployTx.ContractAddress)
		assert.Equal(t, transaction.ClassHash, deployTx.ClassHash)
		assert.Equal(t, transaction.ConstructorCallData, deployTx.ConstructorCallData)
		assert.Equal(t, transaction.Version, deployTx.Version)
	})

	t.Run("deploy account transaction", func(t *testing.T) {
		deployJson, err := os.ReadFile("../.." +
			"/clients/feeder/testdata/mainnet/transaction" +
			"/0xd61fc89f4d1dc4dc90a014957d655d38abffd47ecea8e3fa762e3160f155f2.json")
		assert.NoError(t, err)

		response := new(feeder.TransactionStatus)
		err = json.Unmarshal(deployJson, response)
		require.NoError(t, err)

		transaction := response.Transaction
		txn, err := adaptTransaction(transaction)
		deployAccountTx, ok := txn.(*core.DeployAccountTransaction)
		require.True(t, ok)
		require.NoError(t, err)

		assert.Equal(t, transaction.Hash, deployAccountTx.Hash())
		assert.Equal(t, transaction.ContractAddressSalt, deployAccountTx.ContractAddressSalt)
		assert.Equal(t, transaction.ContractAddress, deployAccountTx.ContractAddress)
		assert.Equal(t, transaction.ClassHash, deployAccountTx.ClassHash)
		assert.Equal(t, transaction.ConstructorCallData, deployAccountTx.ConstructorCallData)
		assert.Equal(t, transaction.Version, deployAccountTx.Version)
		assert.Equal(t, transaction.MaxFee, deployAccountTx.MaxFee)
		assert.Equal(t, transaction.Signature, deployAccountTx.Signature())
		assert.Equal(t, transaction.Nonce, deployAccountTx.Nonce)
	})

	t.Run("declare transaction", func(t *testing.T) {
		declareJson, err := os.ReadFile("../../clients/feeder/testdata/goerli/transaction/0x6eab8252abfc9bbfd72c8d592dde4018d07ce467c5ce922519d7142fcab203f.json")
		assert.NoError(t, err)

		response := new(feeder.TransactionStatus)
		err = json.Unmarshal(declareJson, response)
		require.NoError(t, err)

		transaction := response.Transaction
		txn, err := adaptTransaction(transaction)
		declareTx, ok := txn.(*core.DeclareTransaction)
		require.True(t, ok)
		require.NoError(t, err)

		assert.Equal(t, transaction.Hash, declareTx.Hash())
		assert.Equal(t, transaction.SenderAddress, declareTx.SenderAddress)
		assert.Equal(t, transaction.Version, declareTx.Version)
		assert.Equal(t, transaction.Nonce, declareTx.Nonce)
		assert.Equal(t, transaction.MaxFee, declareTx.MaxFee)
		assert.Equal(t, transaction.Signature, declareTx.Signature())
		assert.Equal(t, transaction.ClassHash, declareTx.ClassHash)
	})

	t.Run("l1handler transaction", func(t *testing.T) {
		deployJson, err := os.ReadFile("../.." +
			"/clients/feeder/testdata/mainnet/transaction" +
			"/0x537eacfd3c49166eec905daff61ff7feef9c133a049ea2135cb94eec840a4a8.json")
		assert.NoError(t, err)

		response := new(feeder.TransactionStatus)
		err = json.Unmarshal(deployJson, response)
		require.NoError(t, err)

		transaction := response.Transaction
		txn, err := adaptTransaction(transaction)
		l1HandlerTx, ok := txn.(*core.L1HandlerTransaction)
		require.True(t, ok)
		require.NoError(t, err)

		assert.Equal(t, transaction.Hash, l1HandlerTx.Hash())
		assert.Equal(t, transaction.ContractAddress, l1HandlerTx.ContractAddress)
		assert.Equal(t, transaction.EntryPointSelector, l1HandlerTx.EntryPointSelector)
		assert.Equal(t, transaction.Nonce, l1HandlerTx.Nonce)
		assert.Equal(t, transaction.CallData, l1HandlerTx.CallData)
		assert.Equal(t, transaction.Version, l1HandlerTx.Version)
	})
}