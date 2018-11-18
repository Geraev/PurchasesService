package purchaseRepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"infoservicenordgold/driver"
	"infoservicenordgold/models"
)

type PurchaseRepository struct{}

func (c PurchaseRepository) GetPurchase(db *sqlx.DB, id string) (*models.Purchase, error) {
	var purchase models.Purchase
	if err := db.Get(&purchase, "SELECT * FROM purchase WHERE id = ?", id); err != nil {
		return nil, err
	}
	return &purchase, nil
}

func (c PurchaseRepository) GetPurchases(db *sqlx.DB) (*[]models.Purchase, error) {
	var purchases []models.Purchase
	if err := db.Select(&purchases, "SELECT * FROM purchase"); err != nil {
		return nil, err
	}
	return &purchases, nil
}

func (c PurchaseRepository) CreatePurchase(db *sqlx.DB, purchase models.Purchase) (*models.Purchase, error) {
	var contractCredits []int
	if err := db.Get(&contractCredits, "SELECT credits FROM contract WHERE id = ?", purchase.ContractId); err != nil {
		return nil, fmt.Errorf(models.ResultSet[3])
	}

	var sumSpent []int
	if err := db.Select(&sumSpent,"SELECT COALESCE(SUM(Spent), 0) as SumSpent FROM purchase WHERE ContractID = ?", purchase.ContractId); err != nil {
		return nil, err
	}

	if sumSpent[0] + purchase.Spent <= contractCredits[0]  {
		purchase.ID = driver.GetUUID()
		nstmt, err := db.PrepareNamed("INSERT INTO purchase VALUES (:ID, :Spent, :ContractId, :DatePursh)")
		if err != nil {
			panic(err)
		}
		nstmt.MustExec(purchase)
		return  &purchase, nil
	}
	return nil, fmt.Errorf(models.ResultSet[5])

}
