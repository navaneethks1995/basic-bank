package db

import (
	"context"
	"testing"
	"time"

	"github.com/navaneethks1995/simplebank/util"
	"github.com/stretchr/testify/require"
)

// createRandomEntry creates a new random Entry and returns it.
// The accountID of the Entry is set to the ID of the given account.
// The amount of the Entry is set to a random money value using util.RandomMoney().
// The function requires there to be no error when creating the Entry,
// and that the Entry is not empty.
// The function also requires that the AccountID and Amount of the created Entry
// are equal to the given accountID and the randomly generated amount respectively.
// The function requires that the ID and CreatedAt fields of the created Entry
// are not zero.
func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

// TestCreateEntry tests the CreateEntry method of the Queries struct.
// It tests that the AccountID, Amount, ID, and CreatedAt fields are properly set.
func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

// TestGetEntry tests the GetEntry method of the Queries struct.
// It tests that the ID, AccountID, Amount, and CreatedAt fields are properly set.
func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)

	entry1 := createRandomEntry(t, acc)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt.Time, entry2.CreatedAt.Time, time.Second)
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesParams{
		AccountID: acc.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
