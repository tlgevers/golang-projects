package types

import (
	"time"
)

type RetailTransaction struct {
	Id       int       `bigquery:"id"`
	DateTime time.Time `bigquery:"datetime"`
	Store    int       `bigquery:"store"`
	Item     string    `bigquery:"item"`
	Amount   float32   `bigquery:"amount"`
	Count    int       `bigquery:"count"`
	Tax      float32   `bigquery:"tax"`
	Cost     float32   `bigquery:"cost"`
}
