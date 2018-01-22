package manipulador

import (
	"net/http"
	"fmt"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"github.com/edias15/api-go/repo"
)

//Deletecompany is function to include new company throught json struct
func Deletecompany(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	
	var countrows int

	id, err := strconv.Atoi(param.ByName("id"))
	if err != nil {
		http.Error(w, "Impossible to read company id.", http.StatusInternalServerError)
		fmt.Println("[deletecompany] Impossible to read parameter id: ", "Erro: ", err.Error())
		return
	}

	err = nil
	sql := "select count(*) as countrows from companies where companyid = " + param.ByName("id")
	rows, err := repo.Db.Query(sql)
	if err != nil {
		http.Error(w, "Not possible reading query.", http.StatusInternalServerError)
		fmt.Println("[deletecompany] Not possible reading query. Erro: ", countrows, err.Error())
		return
	}
	for rows.Next() {
		err = rows.Scan(&countrows)
	}

	/* Didn't find company id in the table database */
	if countrows == 0 {
		http.Error(w, "Didn't find company id in the table database.", http.StatusInternalServerError)
		fmt.Println("[deletecompany] Didn't find company id in the table database")
	} else {
		sql = "Delete from companies where companyid = " + param.ByName("id")
		_, err := repo.Db.Exec(sql)
		if err != nil {
			http.Error(w, "Not possible delete line from database.", http.StatusInternalServerError)
			fmt.Println("[deletecompany] Not possible delete company id: ", id,"Erro: ", err.Error())
			return
		}
		fmt.Println("[deletecompany] Record deleted company id: ", id,"Erro: ", sql)
	}
}
