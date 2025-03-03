package duckdb

import (
	"context"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/query"
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

func TestPointInPolygon(t *testing.T) {

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

	lon := -122.4204643
	lat := 37.7586708

	expected := 4

	pt := orb.Point([2]float64{lon, lat})
	geom := geojson.NewGeometry(pt)

	req := &query.SpatialRequest{
		Geometry: geom,
	}

	q, err := query.NewPointInPolygonQuery(ctx, "pip://")

	if err != nil {
		t.Fatalf("Failed to create new PIP query, %v", err)
	}

	spr, err := query.ExecuteQuery(ctx, db, q, req)

	if err != nil {
		t.Fatalf("Failed to execute PIP query, %v", err)
	}

	results := spr.Results()
	count := len(results)

	if count != expected {
		t.Fatalf("Invalid count (%d), expected %d", count, expected)
	}

	/*
		for _, r := range spr.Results() {
			slog.Info("WTF", "i", r.Id())
		}
	*/
}

func TestIntersects(t *testing.T) {

	t.Skip()
}
