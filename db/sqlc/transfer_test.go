package db

import (
	"context"
	"testing"
	"time"

	"github.com/navaneethks1995/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, acc1, acc2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

// TestCreateTransfer tests the CreateTransfer method of the Queries struct.
// It tests that the FromAccountID, ToAccountID, Amount, and CreatedAt fields are properly set.
func TestCreateTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)
	createRandomTransfer(t, acc1, acc2)
}

// TestGetTransfer tests the GetTransfer method of the Queries struct.
// It tests that the ID, FromAccountID, ToAccountID, and CreatedAt fields are properly set.
func TestGetTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	transfer1 := createRandomTransfer(t, acc1, acc2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.WithinDuration(t, transfer1.CreatedAt.Time, transfer2.CreatedAt.Time, time.Second)
}

// TestListTransfers tests the ListTransfers method of the Queries struct.
// It tests that the ListTransfers method returns a slice of Transfers with the
// correct length and that each Transfer is not empty.

func TestListTransfers(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, acc1, acc2)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID:   acc2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromAccountID == acc1.ID && transfer.ToAccountID == acc2.ID)
	}
}
