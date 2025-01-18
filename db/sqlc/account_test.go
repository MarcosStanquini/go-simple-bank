package db

import (
	"context"
	"testing"

	"github.com/MarcosStanquini/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

// Utilizamos o parametro para gerenciar o  Estado do teste(é um objeto)
// TestCreateAccount testa a criação de uma conta
func TestCreateAccount(t *testing.T) {
	// Struct para representar os parâmetros de entrada
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	//Utilizamos o testQueries que aponta para Queries que é uma struct do SQLC para fazer uso/conexão com o banco.
	//O contextBackground é utilizado para gerenciar tempo de execução ou cancelamento.
	account, err := testQueries.CreateAccount(context.Background(), arg)
	//Verifica se o erro for nulo, caso não for ele falhará o teste automatico(usando a biblioteca, testify)
	//Puxamos o argumento da account na model
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	
}
