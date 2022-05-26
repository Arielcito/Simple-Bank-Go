package db

import (
	"context"
	"testing"

	"github.com/Arielcito/simple-bank-go/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}
func TestGetTransfer(t *testing.T) {
	transfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), int64(transfer1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
}

// func TestUpdateTransfer(t *testing.T) {
// 	transfer1 := CreateRandomTransfer(t)

// 	arg := UpdateTransferParams{
// 		ID:     int64(transfer1.ID),
// 		Amount: util.RandomMoney(),
// 	}
// 	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, transfer2)

// 	require.Equal(t, transfer1.ID, transfer2.ID)
// 	require.Equal(t, arg.ID, transfer2.ID)
// 	require.Equal(t, transfer1.Amount, transfer2.Amount)
// }
// func TestDeleteTransfer(t *testing.T) {
// 	transfer1 := CreateRandomTransfer(t)

// 	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
// 	require.NoError(t, err)

// 	transfer2, err := testQueries.GetEntry(context.Background(), transfer1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, transfer2)
// }

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
