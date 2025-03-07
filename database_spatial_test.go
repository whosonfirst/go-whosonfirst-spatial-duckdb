package duckdb

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/query"
)

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

	spatial_q := &query.SpatialQuery{
		Geometry: geom,
	}

	spatial_fn, err := query.NewSpatialFunction(ctx, "pip://")

	if err != nil {
		t.Fatalf("Failed to create new PIP function, %v", err)
	}

	spr, err := query.ExecuteQuery(ctx, db, spatial_fn, spatial_q)

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

	f_rel_path := "fixtures/1108830809.geojson"
	f_abs_path, err := filepath.Abs(f_rel_path)

	if err != nil {
		t.Fatalf("Failed to derive absolute path for %s, %v", f_rel_path, err)
	}

	r, err := os.Open(f_abs_path)

	if err != nil {
		t.Fatalf("Failed to open %s for reading, %v", f_abs_path, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", f_abs_path, err)
	}

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		t.Fatalf("Failed to unmarshal %s, %v", f_abs_path, err)
	}

	expected := 23

	orb_geom := f.Geometry
	geojson_geom := geojson.NewGeometry(orb_geom)

	spatial_q := &query.SpatialQuery{
		Geometry: geojson_geom,
	}

	spatial_fn, err := query.NewSpatialFunction(ctx, "intersects://")

	if err != nil {
		t.Fatalf("Failed to create new intersects function, %v", err)
	}

	spr, err := query.ExecuteQuery(ctx, db, spatial_fn, spatial_q)

	if err != nil {
		t.Fatalf("Failed to execute intersects query, %v", err)
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

