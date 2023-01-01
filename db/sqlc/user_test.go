package db

import (
	"context"
	"testing"
	"time"

	"github.com/afushimi-source/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

// func TestUpdateuser(t *testing.T) {
// 	user1 := createRandomCreateuser(t)

// 	arg := UpdateuserParams{
// 		ID:      user1.ID,
// 		Balance: util.RandomMoney(),
// 	}

// 	user2, err := testQueries.Updateuser(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, user2)

// 	require.Equal(t, user1.ID, user2.ID)
// 	require.Equal(t, user1.Owner, user2.Owner)
// 	require.Equal(t, arg.Balance, user2.Balance)
// 	require.Equal(t, user1.Currency, user2.Currency)
// 	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
// }

// func TestDeleteuser(t *testing.T) {
// 	user1 := createRandomCreateuser(t)
// 	err := testQueries.Deleteuser(context.Background(), user1.ID)
// 	require.NoError(t, err)

// 	user2, err := testQueries.Getuser(context.Background(), user1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, user2)
// }

// func TestListusers(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomCreateuser(t)
// 	}

// 	arg := ListusersParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	users, err := testQueries.Listusers(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, users)

// 	for _, user := range users {
// 		require.NotEmpty(t, user)
// 	}
// }
