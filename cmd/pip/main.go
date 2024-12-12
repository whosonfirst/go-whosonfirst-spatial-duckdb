package main

import (
	"context"
	"log"

	_ "github.com/whosonfirst/go-whosonfirst-spatial-duckdb"

	"github.com/whosonfirst/go-whosonfirst-spatial/app/pip"
)

func main() {

	ctx := context.Background()
	err := pip.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
