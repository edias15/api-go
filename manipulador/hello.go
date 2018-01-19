package manipulador

import (
	"fmt"
	"net/http"
	"time"

	"github.com/edias15/eXperto/model"
)

//Hello is the manipulator of route requisition
func Hello(w http.ResponseWriter, r *http.Request) {
	horas := time.Now().Format("15:04:05")
	page := model.Page{}
	page.Hora = horas
	if err := HelloModel.ExecuteTemplate(w,"hello.html",page); err != nil {
		http.Error(w, "There is a trouble with page renderization.", http.StatusInternalServerError)
		fmt.Println("[Hello] Error on execute model: ",err.Error())
	}
}