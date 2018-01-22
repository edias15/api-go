package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/edias15/api-go/manipulador"
	"github.com/edias15/api-go/repo"
)

func init() {
	fmt.Println("Server uploading...")
}

func main() {
	err := repo.AbreConexaoComBDPostgres()
	if err != nil {
		fmt.Println("Error on open database : ",err.Error())
	}
	router := httprouter.New()
	router.GET("/",manipulador.Hello)
	router.GET("/company/:id",manipulador.Company)
	router.POST("/insertcompany/:id/:name/:cnpj",manipulador.Insertcompany)
//	router.DELETE("/deletecompany/:id",manipulador.Deletecompany)

//	http.HandleFunc("/", manipulador.Hello)
//	http.HandleFunc("/company/", manipulador.Company)
	fmt.Println("Estou ouvindo...")
	http.ListenAndServe(":8181", router)
}
