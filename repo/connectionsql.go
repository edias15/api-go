package repo

import (
	"github.com/jmoiron/sqlx"
	/* github.com/lib/pq não é usado direto pela aplicação */
	_ "github.com/lib/pq"
)
//Db armazena a conexão com o banco  de dados
var Db *sqlx.DB

//AbreConexaoComBDPostgres função que abre conexão com o BD 
func AbreConexaoComBDPostgres() (err error) {
	err = nil
	Db, err = sqlx.Open("postgres","postgres://tdfqgrqkpnitww:3a6a2b3fae7c6988408d696860539a2021bade6e04069d47d84200cee8335b5b@ec2-107-20-193-202.compute-1.amazonaws.com:5432/d907s0p5lek9bb")
	if err != nil {
		return
	}

	err = Db.Ping()
	if err != nil {
		return
	}
	defer Db.Close()

	return
}