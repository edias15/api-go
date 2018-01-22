package model

//CompaniesFormat format of table: companies
type CompaniesFormat struct {
	Codempresa int `json:"companyid" db:"companyid"`
	Nome string `json:"name" db:"name"`
	Cnpjs string `json:"cnpj" db:"cnpj"`
}
