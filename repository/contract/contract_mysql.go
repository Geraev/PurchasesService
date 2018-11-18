package contractRepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"infoservicenordgold/driver"
	"infoservicenordgold/models"
)

type ContractRepository struct{}

func (c ContractRepository) GetContract(db *sqlx.DB, id string) (*models.Contract, error) {
	var contract models.Contract
	if err := db.Get(&contract, "SELECT * FROM contract WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &contract, nil
}

func (c ContractRepository) GetContracts(db *sqlx.DB) (*[]models.Contract, error) {
	var contracts []models.Contract
	if err := db.Select(&contracts, "SELECT * FROM contract"); err != nil {
		return nil, err
	}
	return &contracts, nil
}

func (c ContractRepository) CreateContract(db *sqlx.DB, contract models.Contract) (*models.Contract, error) {
	contract.ID = driver.GetUUID()
	nstmt, err := db.PrepareNamed("INSERT INTO contract VALUES (:ID, :SellerId, :ClientId, :ContractNumber, :SignedDate, :ValidTill, :Credits)")
	if err != nil {
		return nil, err
	}
	nstmt.MustExec(contract)
	return &contract, nil
}

func (c ContractRepository) UpdateContract(db *sqlx.DB, contract models.Contract) (*models.Contract, error) {
	nstmt, err := db.PrepareNamed("UPDATE contract SET SellerId = :SellerId, ClientId = :ClientId, ContractNumber = :ContractNumber, SignedDate = :SignedDate, ValidTill = :ValidTill, Credits = :Credits WHERE id = :ID")
	if err != nil {
		return nil, err
	}
	if rowsAff, err := nstmt.MustExec(contract).RowsAffected(); err != nil {
		return nil, err
	} else if rowsAff == 0 {
		return nil, fmt.Errorf("ID not found: ", contract)
	}
	return &contract, nil
}
