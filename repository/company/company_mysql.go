package companyRepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"infoservicenordgold/driver"
	"infoservicenordgold/models"
)

type CompanyRepository struct{}

func (c CompanyRepository) GetCompany(db *sqlx.DB, id string) (*models.Company, error) {
	var company models.Company
	if err := db.Get(&company, "SELECT * FROM company WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &company, nil
}

func (c CompanyRepository) GetCompanies(db *sqlx.DB) (*[]models.Company, error) {
	var companies []models.Company
	if err := db.Select(&companies, "SELECT * FROM company"); err != nil {
		return nil, err
	}
	return &companies, nil
}

func (c CompanyRepository) CreateCompany(db *sqlx.DB, company models.Company) (*models.Company, error) {
	company.ID = driver.GetUUID()
	nstmt, err := db.PrepareNamed("INSERT INTO Company VALUES (:ID, :CompanyName, :RegistrationCode)")
	if err != nil {
		return nil, err
	}
	nstmt.MustExec(company)
	return &company, nil
}

func (c CompanyRepository) UpdateCompany(db *sqlx.DB, company models.Company) (*models.Company, error) {
	nstmt, err := db.PrepareNamed("UPDATE Company SET CompanyName = :CompanyName, RegistrationCode = :RegistrationCode WHERE id = :ID")
	if err != nil {
		return nil, err
	}
	if rowsAff, err := nstmt.MustExec(company).RowsAffected(); err != nil {
		return nil, err
	} else if rowsAff == 0 {
		return nil, fmt.Errorf("ID not found: ", company)
	}
	return &company, nil
}

