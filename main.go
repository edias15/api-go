package main

import (
	"fmt"
	"net/http"
	"github.com/edias15/eXperto/manipulador"
	"github.com/edias15/eXperto/repo"
)

func init() {
	fmt.Println("Server uploading...")
}

func main() {
	err := repo.AbreConexaoComBDPostgres()
	if err != nil {
		fmt.Println("Error on open database : ",err.Error())
	}
	http.HandleFunc("/", manipulador.Hello)
	http.HandleFunc("/company/", manipulador.Company)
	fmt.Println("Estou ouvindo...")
	http.ListenAndServe(":8181", nil)
}
