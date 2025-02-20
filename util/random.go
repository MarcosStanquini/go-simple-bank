package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Criando um novo gerador de números aleatórios
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomInt(min, max int64) int64 {
	return min + r.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c) // Aqui, como os caracteres são ASCII, podemos usar WriteByte diretamente
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "BRL", "CSD"}
	n := len(currencies)
	return currencies[r.Intn(n)] // Usando 'r' para o gerador de números aleatórios
}
