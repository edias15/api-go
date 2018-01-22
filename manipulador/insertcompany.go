package manipulador

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/edias15/api-go/repo"
  	"github.com/edias15/api-go/model"
)

//Insertcompany is function to include new company throught json struct
func Insertcompany(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	companies := model.CompaniesFormat{}

	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		http.Error(w, "Impossible to read company id.", http.StatusInternalServerError)
		fmt.Println("[insertcompany] Impossible to read parameter id: ", "Erro: ", err.Error())
		return
	}
	name := param.ByName("name")
	cnpj := param.ByName("cnpj")

	sql := "select companyid, name, cnpj from companies where companyid = $1"
	line, err := repo.Db.Queryx(sql, id)
	if err != nil {
		http.Error(w, "Impossible search this number.", http.StatusInternalServerError)
		fmt.Println("[insertcompany] Not possible executing query: ", sql, "Erro: ", err.Error())
		return
	}
	for line.Next() {
		err = line.StructScan(&companies)
		if err != nil {
			http.Error(w, "Does not possible searching this number",http.StatusInternalServerError)
			fmt.Println("[insertcompany] Not possible binding data in the struct. Erro: ", err.Error())
			return
		}
	}
	/* if query no results then Codempresa = 0 */
	if companies.Codempresa != 0 {
		http.Error(w, "The company id exist in database.", http.StatusInternalServerError)
		fmt.Println("[insertcompany] The company id exist in database id: ", id)
	} else {
		data := map[string]interface{}{
            "id": id,
            "name": name,
			"cnpj": cnpj,
		}
		_, err := repo.Db.NamedExec(`insert into companies (companyid, name, cnpj) values (:id, :name, :cnpj)`, data)
		if err != nil {
			http.Error(w, "Not possible insert new line into database.", http.StatusInternalServerError)
			fmt.Println("[insertcompany] Not possible insert new line into database: ", data, "Erro: ", err.Error())
			return
		}
	}
}
