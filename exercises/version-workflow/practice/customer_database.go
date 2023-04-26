package loanprocess

import (
	"errors"
)

type CustomerInfo struct {
	CustomerID      string
	Name            string
	EmailAddress    string
	Amount          int
	NumberOfPeriods int
}

type CustomerInfoDatabase interface {
	Get(customerID string) (CustomerInfo, error)
}

var customers map[string]CustomerInfo

type simpleCustomerMap struct {
}

func CustomerInfoDB() CustomerInfoDatabase {
	db := simpleCustomerMap{}
	db.populate()

	return db
}

func (db *simpleCustomerMap) populate() {
	customers = make(map[string]CustomerInfo)

	customer01 := CustomerInfo{
		CustomerID:      "a100",
		Name:            "Ana Garcia",
		EmailAddress:    "ana@example.com",
		Amount:          500,
		NumberOfPeriods: 10,
	}

	customer02 := CustomerInfo{
		CustomerID:      "a101",
		Name:            "Amit Singh",
		EmailAddress:    "asingh@example.com",
		Amount:          250,
		NumberOfPeriods: 15,
	}

	customer03 := CustomerInfo{
		CustomerID:      "a102",
		Name:            "Mary O'Connor",
		EmailAddress:    "marymo@example.com",
		Amount:          425,
		NumberOfPeriods: 12,
	}

	customers[customer01.CustomerID] = customer01
	customers[customer02.CustomerID] = customer02
	customers[customer03.CustomerID] = customer03
}

func (db simpleCustomerMap) Get(customerID string) (CustomerInfo, error) {
	info, defined := customers[customerID]
	if !defined {
		return CustomerInfo{}, errors.New("customer ID does not exist in the database")
	}

	return info, nil
}
