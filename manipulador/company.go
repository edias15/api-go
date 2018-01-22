package manipulador

import (
	"strconv"
	"net/http"
	"fmt"

	"github.com/julienschmidt/httprouter"
	"github.com/edias15/api-go/repo"
  	"github.com/edias15/api-go/model"
)

//Company is function to manipulate route requisition
func Company(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	companies := model.CompaniesFormat{}
	companyid, err := strconv.Atoi(param.ByName("id"))

	if err != nil {
		http.Error(w, "It wasn't typed a number for company id.", http.StatusBadRequest)
		fmt.Println("[company] Error converting string to number:", err.Error())
	}

	sql := "select companyid, name, cnpj from companies where companyid = $1"
	line, err := repo.Db.Queryx(sql, companyid)
	if err != nil {
		http.Error(w, "Impossible search this number.", http.StatusInternalServerError)
		fmt.Println("[company] Not possible executing query: ", sql, "Erro: ", err.Error())
		return
	}
	for line.Next() {
		err = line.StructScan(&companies)

		if err != nil {
			http.Error(w, "Does not possible searching this number",http.StatusInternalServerError)
			fmt.Println("[company] Not possible binding data in the struct. Erro: ", err.Error())
			return
		}
	}
	if err := CompanyModel.ExecuteTemplate(w,"company.html",companies); err != nil {
		http.Error(w, "Error in page renderization.", http.StatusInternalServerError)
		fmt.Println("[company] Error executing model", err.Error())
		return
	}
}
