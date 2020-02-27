package main

import (
	"fmt"
	"github.com/tlgevers/golang-projects/random-sales/pkg/types"
	bq "github.com/tlgevers/golang-projects/random-sales/pkg/bigquery"
	"math/rand"
	"os"
	"time"
)

var (
	rows       = 1000
	project_id = os.Getenv("PROJECT_ID")
	depts      = [...]string{"Coffee", "Mug", "Beans", "Sticker", "T-Shirt", "Flavor"}
)

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomDate() (t time.Time) {
	year := randomInt(2010, 2020)
	month := randomInt(1, 12)
	day := randomInt(1, 28)
	hour := randomInt(7, 20)
	min := randomInt(1, 60)
	sec := randomInt(1, 60)

	t = time.Date(year, time.Month(month), day, hour, min, sec, 651387237, time.UTC)
	return
}

func randomSale() types.RetailTransaction {
	fmt.Println("Creating Sale")
	rand.Seed(time.Now().UnixNano())
	id := time.Now().UnixNano()
	dt := randomDate()
	store := randomInt(1, 2639)
	dept := depts[rand.Intn(6)]
	amount := 3 + rand.Float32()*(20-3)
	count := rand.Intn(6-1) + 1
	cost := amount * 0.5
	transaction := types.RetailTransaction{
		Id:       int(id),
		DateTime: dt,
		Store:    store,
		Item:     dept,
		Amount:   amount,
		Count:    count,
		Tax:      amount * 0.1,
		Cost:     cost,
	}
	return transaction
}

func main() {
	//bq.CreateTable(project_id)
	start := time.Now()
	fmt.Println("Generating Random Sales...")
	var transactions []types.RetailTransaction
	for i := 0; i < rows; i++ {
		sale := randomSale()
		fmt.Printf("sale: %+v", sale)
		transactions = append(transactions, sale)
	}
	bq.InsertRows(project_id, "kss", "magicbeans_transactions", transactions)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Time elapsed", elapsed, "seconds")
}
