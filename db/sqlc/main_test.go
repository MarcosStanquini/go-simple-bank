package db

import(
	 "testing"
	 "log"
	 "database/sql"
	 "os"
	 _ "github.com/lib/pq"
)

const(
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil{
		log.Fatal("Não pode ser conectado a database:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
	 
}