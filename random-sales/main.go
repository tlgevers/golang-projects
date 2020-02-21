package main

import (
	"fmt"
	bq "github.com/tlgevers/golang-projects/random-sales/pkg/bigquery"
	types "github.com/tlgevers/golang-projects/random-sales/pkg/types"
	"math/rand"
	"os"
	"time"
)

var (
	rows       = 10000
	project_id = os.Getenv("PROJECT_ID")
	depts      = [...]string{"Coffee", "Mug", "Beans", "Sticker", "T-Shirt", "Flavor"}
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomSale() types.RetailTransaction {
	fmt.Println("Creating Sale")
	rand.Seed(time.Now().UnixNano())
	id := time.Now().UnixNano()
	store := randomInt(1, 2639)
	dept := depts[rand.Intn(6)]
	amount := 3 + rand.Float32()*(20-3)
	count := rand.Intn(6-1) + 1
	cost := amount * 0.5
	transaction := types.RetailTransaction{
		Id:     int(id),
		Store:  store,
		Item:   dept,
		Amount: amount,
		Count:  count,
		Tax:    amount * 0.1,
		Cost:   cost,
	}
	return transaction
}

func main() {
	start := time.Now()
	fmt.Println("Generating Random Sales...")
	var transactions []types.RetailTransaction
	for i := 0; i < rows; i++ {
		transactions = append(transactions, randomSale())
	}
	bq.InsertRows(project_id, "education", "transactions_magicbeans", transactions)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Time elapsed", elapsed, "seconds")
}
