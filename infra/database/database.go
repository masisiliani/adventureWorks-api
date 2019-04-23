package database

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

//Connect: Inicializa a conexão com a base de dados e retorna uma struct do tipo DB, que tem os valores para manipulação do banco de dados
func Connect() error {

	//Declaração da variável [stringConnection] que vai ser preenchida com os valores de acesso da base
	var stringConnection = "sqlserver://msisiliani:Admin123@devwoman.database.windows.net?database=AdventureWorksLT&app name=adventureWorks-api"

	//Abre a conexão com a base, recebendo como retorno a conexão e um tratativa de erro
	db, err := sql.Open("mssql", stringConnection)

	//Verificação de erro do retorno da abertura da conexão
	if err != nil {
		return err
	}

	//Esse trecho verifica se a conxão está aberta e ativa.
	err = db.Ping()
	//Verificação de erro
	if err != nil {
		return err
	}

	DB = db

	return nil
}
