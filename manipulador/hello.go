package manipulador

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	"time"

	"github.com/edias15/api-go/model"
)

//Hello is the manipulator of route requisition
func Hello(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	horas := time.Now().Format("15:04:05")
	page := model.Page{}
	page.Hora = horas
	if err := HelloModel.ExecuteTemplate(w,"hello.html",page); err != nil {
		http.Error(w, "There is a trouble with page renderization.", http.StatusInternalServerError)
		fmt.Println("[Hello] Error on execute model: ",err.Error())
	}
}