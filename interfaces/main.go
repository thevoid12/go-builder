// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.
// There are all sorts of reasons that you might end up using a interface in Go, but in my experience the three most common are:
// To help reduce duplication or boilerplate code.
// To make it easier to use mocks instead of real objects in unit tests.
// As an architectural tool, to help enforce decoupling between parts of your codebase.
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// Create our own custom ShopModel interface. Notice that it is perfectly
// fine for an interface to describe multiple methods, and that it should
// describe input parameter types as well as return value types.
type ShopModel interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (float64, error)
}

// The ShopDB type satisfies our new custom ShopModel interface, because it
// has the two necessary methods -- CountCustomers() and CountSales().
type ShopDB struct {
	*sql.DB
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (float64, error) {
	var count float64
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func main() {
	db, err := sql.Open("postgres", "postgres://user:pass@localhost/db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	shopDB := &ShopDB{db}
	sr, err := calculateSalesRate(shopDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(sr)
}

// Swap this to use the ShopModel interface type as the parameter, instead of the
// concrete *ShopDB type.
func calculateSalesRate(sm ShopModel) (string, error) {
	since := time.Now().Add(-24 * time.Hour)

	sales, err := sm.CountSales(since)
	if err != nil {
		return "", err
	}

	customers, err := sm.CountCustomers(since)
	if err != nil {
		return "", err
	}

	rate := float64(sales) / float64(customers)
	return fmt.Sprintf("%.2f", rate), nil
}
