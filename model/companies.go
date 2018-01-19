package model

//Companies Table of Companies
type Companies struct {
	Codempresa int `json:"codempresa" db:"companyid"`
	Nome string `json:"nome" db:"name"`
	Cnpjs string `json:"cnpjs" db:"cnpj"`
}
