package models

import "time"

type Company struct {
	ID      string `db:"ID" json:"Id"`
	Name    string `db:"CompanyName" json:"CompanyName"`
	RegCode string `db:"RegistrationCode" json:"RegistrationCode"`
}

type Contract struct {
	ID          string    `db:"ID" json:"Id"`
	Seller      string    `db:"SellerId" json:"sellerId"`
	Client      string    `db:"ClientId" json:"clientId"`
	NumContract string    `db:"ContractNumber" json:"contractNumber"`
	SignedDate  time.Time `db:"SignedDate" json:"signedDate"`
	ValidTill   time.Time `db:"ValidTill" json:"validTill"`
	Credits     int       `db:"Credits" json:"Credits"`
}

type Purchase struct {
	ID         string    `db:"ID" json:"Id"`
	Spent      int       `db:"Spent" json:"Spent"`
	ContractId string    `db:"ContractId" json:"ContractId"`
	DatePursh  time.Time `db:"DatePursh" json:"DatePursh"`
}

type RestResult struct {
	ResultCode   int    `json:"resultCode"`
	ResultStatus string `json:"resultStatus"`
}

var ResultSet = map[int]string{
	1: "Delete operation completed successfully",
	2: "Company with this id does not exist",
	3: "Contract with this id does not exist",
	4: "Purchase with this id does not exist",
	5: "Общая сумма покупок превышает стоимость контракта",
}
