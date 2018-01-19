
package manipulador

import "html/template"

//HelloModel armazena o modelo de pagina hello
var HelloModel = template.Must(template.ParseFiles("view/hello.html"))

//CompanyModel armazena o modelos de pagina company
var CompanyModel = template.Must(template.ParseFiles("view/company.html"))