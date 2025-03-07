package duckdb

import (
	"context"
	"fmt"
	"io"
	"path/filepath"
	"testing"

	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
)

func TestReader(t *testing.T) {

	t.Skip()

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

	defer db.Disconnect(ctx)

	f_uri := "1108831819.geojson"
	
	r, err := db.Read(ctx, f_uri)

	if err != nil {
		t.Fatalf("Failed to open %s for reading, %v", f_uri, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", f_uri, err)
	}

	_, err = geojson.UnmarshalFeature(body)

	if err != nil {
		t.Fatalf("Failed to unmarshal %s as GeoJSON, %v", f_uri, err)
	}
}

