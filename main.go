package main

import (
	"github.com/jmoiron/sqlx"
	"infoservicenordgold/controllers"
	"infoservicenordgold/driver"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var db *sqlx.DB

func init() {
	gotenv.Load()
}

type WithCORS struct {
	r *mux.Router
}

// Wrapper to Allow CORS.
func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(res, req)
}

func main() {
	db = driver.ConnectDB()
	defer db.Close()

	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/companies", controller.GetCompanies(db)).Methods("GET")
	router.HandleFunc("/company/{id}", controller.GetCompany(db)).Methods("GET")
	router.HandleFunc("/company", controller.CreateCompany(db)).Methods("POST")
	router.HandleFunc("/company", controller.UpdateCompany(db)).Methods("PUT")

	router.HandleFunc("/contracts", controller.GetContracts(db)).Methods("GET")
	router.HandleFunc("/contract/{id}", controller.GetContract(db)).Methods("GET")
	router.HandleFunc("/contract", controller.CreateContract(db)).Methods("POST")
	router.HandleFunc("/contract", controller.UpdateContract(db)).Methods("PUT")

	router.HandleFunc("/purchases", controller.GetPurchases(db)).Methods("GET")
	router.HandleFunc("/purchase/{id}", controller.GetPurchase(db)).Methods("GET")
	router.HandleFunc("/purchase", controller.CreatePurchase(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", &WithCORS{router}))
}
