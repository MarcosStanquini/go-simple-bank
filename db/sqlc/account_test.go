package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/MarcosStanquini/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func creatRandomAccount(t *testing.T) Account {
	// Struct para representar os parâmetros de entrada
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	//Utilizamos o testQueries que aponta para Queries que é uma struct do SQLC para fazer uso/conexão(uma instancia) com o banco.
	//O contextBackground(principalmente utilizado pelo sqlc) é utilizado para gerenciar tempo de execução ou cancelamento.
	//Passo o arg que é a conta gerada com dados aleatórios para a criação da conta
	account, err := testQueries.CreateAccount(context.Background(), arg)
	//Verifica se o erro for nulo, caso não for ele falhará o teste automatico(usando a biblioteca, testify)
	//Puxamos o argumento da account na model
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner) //Verifico se o arg criado com dados aleatórios é o mesmo que está no banco
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// Utilizamos o parametro para gerenciar o  Estado do teste(é um objeto)
// TestCreateAccount testa a criação de uma conta
func TestCreateAccount(t *testing.T) {
	creatRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account1 := creatRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}


func TestUpdateAccount(t *testing.T){
	account1 := creatRandomAccount(t)

	arg := UpdateAccountParams{
		ID: account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.Equal(t, account1.Owner, account2.Owner)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T){
	account1 := creatRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)


}

func TestListAccount(t *testing.T){
	for i:=0; i < 10; i++{
		creatRandomAccount(t)
	}

	arg := ListAccountsParams{
		 Limit: 5,
		 Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	//To omitindo um valor(i provavelmente, ou seja um blank identifier), passando o identificador para iterar sobre cada conta no range(tamanho da lista de contas que eu criei anteriormente)
	for _, account := range accounts{
		require.NotEmpty(t, account)
	}


}



	 

