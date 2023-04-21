package conflux

import (
	"context"
	"testing"

	"gotest.tools/assert"
)

func TestOnConfluxGetBlockByEpochNumber(t *testing.T) {
	c := NewClient()

	ctx := context.Background()
	// curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"cfx_getBlockByEpochNumber","params":["0x427ced0", true],"id":67}' https://main.confluxrpc.com
	t.Run("test cft_getBlockByEpochNumber", func(t *testing.T) {
		blockInfo, err := c.GetBlockTransactions(ctx, GetBlockTransactionsParameter{
			BlockNumber: 69717712,
		})
		assert.NilError(t, err)
		assert.Equal(t, 1, len(blockInfo.Transactions))
		assert.Equal(t, "0x30bbb39f7d5d74bc1442b92e8ac786e20c139d1135438e4afb2ecc1e77da1a77", blockInfo.Transactions[0].Hash)
		assert.Equal(t, "0x76fb67590afb26d599b9f531c2b31053b02f050180e56a4523cde7e6cdc6f6d7", blockInfo.Transactions[0].BlockHash)
		assert.Equal(t, "0x0", blockInfo.Transactions[0].TransactionIndex)
		assert.Equal(t, "cfx:aajj1b1gm7k51mhzm80czcx31kwxrm2f6jxvy30mvk", blockInfo.Transactions[0].From)
		assert.Equal(t, "cfx:aasjp2fash9hpc6512xbbvc56ymzx60jreb378ff7k", blockInfo.Transactions[0].To)
		assert.Equal(t, "0xce25a8772d19e358", blockInfo.Transactions[0].Value)
		assert.Equal(t, "0x3b9aca00", blockInfo.Transactions[0].GasPrice)
		assert.Equal(t, "0xe4e1c0", blockInfo.Transactions[0].Gas)
		assert.Equal(t, "0x", blockInfo.Transactions[0].Data)
		assert.Equal(t, "0x0", blockInfo.Transactions[0].StorageLimit)
		assert.Equal(t, "0x427cecc", blockInfo.Transactions[0].EpochHeight)
		assert.Equal(t, "0x405", blockInfo.Transactions[0].ChainId)
		assert.Equal(t, "0x0", blockInfo.Transactions[0].Status)
		assert.Equal(t, "0x1", blockInfo.Transactions[0].V)
		assert.Equal(t, "0x146dec897a0f0bd9ba48b9b5bb1594a3b562910e06f95adecab6ecd3909e7aeb", blockInfo.Transactions[0].R)
		assert.Equal(t, "0x3c0aa10a03882d242fb087816e2832e1a8c06850d477297d36e92e4c062d3b38", blockInfo.Transactions[0].S)
		assert.Equal(t, "0x415e2", blockInfo.Transactions[0].Nonce)
	})

	// curl -H "Content-Type: application/json" -X POST --data '{"jsonrpc":"2.0","method":"cfx_getTransactionReceipt","params":["0x30bbb39f7d5d74bc1442b92e8ac786e20c139d1135438e4afb2ecc1e77da1a77"],"id":67}' https://main.confluxrpc.com
	t.Run("test cfx_getTransactionReceipt", func(t *testing.T) {
		receiptInfo, err := c.GetTransactionReceipt(ctx, "0x30bbb39f7d5d74bc1442b92e8ac786e20c139d1135438e4afb2ecc1e77da1a77")
		assert.NilError(t, err)
		assert.Equal(t, "0x30bbb39f7d5d74bc1442b92e8ac786e20c139d1135438e4afb2ecc1e77da1a77", receiptInfo.TransactionHash)
		assert.Equal(t, "0x76fb67590afb26d599b9f531c2b31053b02f050180e56a4523cde7e6cdc6f6d7", receiptInfo.BlockHash)
		assert.Equal(t, "cfx:aajj1b1gm7k51mhzm80czcx31kwxrm2f6jxvy30mvk", receiptInfo.From)
		assert.Equal(t, "cfx:aasjp2fash9hpc6512xbbvc56ymzx60jreb378ff7k", receiptInfo.To)
		assert.Equal(t, "0x5208", receiptInfo.GasUsed)
		assert.Equal(t, "0x27f7d0bdb92000", receiptInfo.GasFee)
		assert.Equal(t, "0x0", receiptInfo.Index)
		assert.Equal(t, "0x427ced0", receiptInfo.EpochNumber)
	})
}

func TestOnConfluxScanGetAccountTx(t *testing.T) {
	c := NewClient()
	ctx := context.Background()
	t.Run("test scan api getAccountTransactions", func(t *testing.T) {
		resp, err := c.GetAccountTransactions(ctx, GetAccountParameter{
			Address: "cfx:aajj1b1gm7k51mhzm80czcx31kwxrm2f6jxvy30mvk",
			Limit:   10,
		})
		assert.NilError(t, err)
		t.Log(resp)
	})

	t.Run("test scan api getAccountTransfers", func(t *testing.T) {
		_, err := c.GetAccountTransfers(ctx, GetAccountParameter{
			Address: "cfx:aarng0xfy95yfxn59t15r58gks20vzn88esp76ag6f",
			Limit:   100,
		})
		assert.NilError(t, err)
	})
}
