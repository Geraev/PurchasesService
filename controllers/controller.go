package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	. "infoservicenordgold/models"
	"infoservicenordgold/repository/company"
	"infoservicenordgold/repository/contract"
	"infoservicenordgold/repository/purchase"
	"net/http"
)

var companyRepo = companyRepository.CompanyRepository{}
var contractRepo = contractRepository.ContractRepository{}
var purchaseRepo = purchaseRepository.PurchaseRepository{}

type Controller struct{}

func (c Controller) GetCompany(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		company, err := companyRepo.GetCompany(db, params["id"])
		if err != nil {
			if err2 := json.NewEncoder(w).Encode(&RestResult{ResultCode: 2, ResultStatus: ResultSet[2]}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(company); err != nil {
			panic(err)
		}
	}
}

func (c Controller) GetCompanies(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		companies, err := companyRepo.GetCompanies(db)
		if err != nil {
			if err2 := json.NewEncoder(w).Encode([]Company{}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(companies); err != nil {
			panic(err)
		}
	}
}

func (c Controller) CreateCompany(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var company Company
		if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
			panic(err)
		}

		if createdCompany, err := companyRepo.CreateCompany(db, company); err != nil {
			panic(err)
		} else if err := json.NewEncoder(w).Encode(createdCompany); err != nil {
			panic(err)
		}
	}
}


func (c Controller) UpdateCompany(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var company Company
		if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
			panic(err)
		}

		if createdCompany, err := companyRepo.UpdateCompany(db, company); err != nil {
			panic(err)
		} else if err := json.NewEncoder(w).Encode(createdCompany); err != nil {
			panic(err)
		}

	}
}



func (c Controller) GetContract(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		contract, err := contractRepo.GetContract(db, params["id"])
		if err != nil {
			if err2 := json.NewEncoder(w).Encode(&RestResult{ResultCode: 2, ResultStatus: ResultSet[2]}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(contract); err != nil {
			panic(err)
		}
	}
}

func (c Controller) GetContracts(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		contracts, err := contractRepo.GetContracts(db)
		if err != nil {
			if err2 := json.NewEncoder(w).Encode([]Contract{}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(contracts); err != nil {
			panic(err)
		}
	}
}

func (c Controller) CreateContract(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var contract Contract
		if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
			panic(err)
		}

		if createdContract, err := contractRepo.CreateContract(db, contract); err != nil {
			panic(err)
		} else if err := json.NewEncoder(w).Encode(createdContract); err != nil {
			panic(err)
		}
	}
}

func (c Controller) UpdateContract(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var contract Contract
		if err := json.NewDecoder(r.Body).Decode(&contract); err != nil {
			panic(err)
		}

		if createdContract, err := contractRepo.UpdateContract(db, contract); err != nil {
			panic(err)
		} else if err := json.NewEncoder(w).Encode(createdContract); err != nil {
			panic(err)
		}

	}
}



func (c Controller) GetPurchase(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		purchase, err := purchaseRepo.GetPurchase(db, params["id"])
		if err != nil {
			if err2 := json.NewEncoder(w).Encode(&RestResult{ResultCode: 2, ResultStatus: ResultSet[2]}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(purchase); err != nil {
			panic(err)
		}
	}
}

func (c Controller) GetPurchases(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		purchases, err := purchaseRepo.GetPurchases(db)
		if err != nil {
			if err2 := json.NewEncoder(w).Encode([]Purchase{}); err2 != nil {
				panic(err2)
			}
		}
		if err := json.NewEncoder(w).Encode(purchases); err != nil {
			panic(err)
		}
	}
}

func (c Controller) CreatePurchase(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var purchase Purchase
		if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
			panic(err)
		}

		if createdPurchase, err := purchaseRepo.CreatePurchase(db, purchase); err != nil {
			panic(err)
		} else if err := json.NewEncoder(w).Encode(createdPurchase); err != nil {
			panic(err)
		}
	}
}
