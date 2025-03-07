package duckdb

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/whosonfirst/go-whosonfirst-spatial/database"
)

func TestDatabase(t *testing.T) {

	ctx := context.Background()

	rel_path := "fixtures/sf_county.parquet"
	abs_path, err := filepath.Abs(rel_path)

	if err != nil {
		t.Fatalf("Failed to derive absolute path for %s, %v", rel_path, err)
	}

	uri := fmt.Sprintf("duckdb://?uri=%s", abs_path)

	db, err := database.NewSpatialDatabase(ctx, uri)

	if err != nil {
		t.Fatalf("Failed to create new database, %v", err)
	}

	err = db.Disconnect(ctx)

	if err != nil {
		t.Fatalf("Failed to disconnect database, %v", err)
	}

}
