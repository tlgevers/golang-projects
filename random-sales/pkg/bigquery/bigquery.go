package bigquery

import (
	bq "cloud.google.com/go/bigquery"
	"context"
	"fmt"
	"github.com/tlgevers/golang-projects/random-sales/pkg/types"
)

func client(projectID string) (client *bq.Client, err error) {
	ctx := context.Background()
	client, err = bq.NewClient(ctx, projectID)
	if err != nil {
		return
	}
	return
}

func InsertRows(projectID string, dataset string, table string, items []types.RetailTransaction) {
	fmt.Println("Inserting ...")
	ctx := context.Background()
	client, err := client(projectID)
	if err != nil {
		panic(err)
	}
	ins := client.Dataset(dataset).Table(table).Inserter()
	if err := ins.Put(ctx, items); err != nil {
		panic(err)
	}
}


