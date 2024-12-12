package duckdb

import (
	"context"
	"database/sql"
	"fmt"
	"iter"
	"log/slog"
	"net/url"
	"time"

	_ "github.com/marcboeker/go-duckdb"

	"github.com/paulmach/orb"
	"github.com/whosonfirst/go-whosonfirst-spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spatial/filter"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
	"github.com/whosonfirst/go-whosonfirst-uri"
)

type DuckDBSpatialDatabase struct {
	database.SpatialDatabase
	conn         *sql.DB
	database_uri string
}

func init() {
	ctx := context.Background()
	database.RegisterSpatialDatabase(ctx, "duckdb", NewDuckDBSpatialDatabase)
}

func NewDuckDBSpatialDatabase(ctx context.Context, uri string) (database.SpatialDatabase, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	q := u.Query()
	database_uri := q.Get("uri")

	conn, err := sql.Open("duckdb", "")

	if err != nil {
		return nil, fmt.Errorf("Failed to create database connection, %w", err)
	}

	extensions := []string{
		"spatial",
		"httpfs",
	}

	for _, ext := range extensions {

		cmds := []string{
			fmt.Sprintf("INSTALL %s", ext),
			fmt.Sprintf("LOAD %s", ext),
		}

		for _, q := range cmds {

			_, err := conn.ExecContext(ctx, q)

			if err != nil {
				return nil, fmt.Errorf("Failed to configure data - query failed, %w (%s)", err, q)
			}
		}
	}

	db := &DuckDBSpatialDatabase{
		conn:         conn,
		database_uri: database_uri,
	}

	return db, nil
}

func (db *DuckDBSpatialDatabase) IndexFeature(context.Context, []byte) error {
	return fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) RemoveFeature(context.Context, string) error {
	return fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) PointInPolygon(ctx context.Context, coord *orb.Point, filters ...spatial.Filter) (spr.StandardPlacesResults, error) {

	results := make([]spr.StandardPlacesResult, 0)
	var err error

	for wof_spr, pip_err := range db.pointInPolygon(ctx, coord, filters...) {

		if err != nil {
			err = pip_err
			break
		}

		results = append(results, wof_spr)
	}

	if err != nil {
		return nil, err
	}

	return NewSPRResults(results), nil
}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidates(ctx context.Context, centroid *orb.Point, filters ...spatial.Filter) ([]*spatial.PointInPolygonCandidate, error) {

	candidates := make([]*spatial.PointInPolygonCandidate, 0)
	var err error

	for wof_spr, pip_err := range db.pointInPolygon(ctx, centroid, filters...) {

		if err != nil {
			err = pip_err
			break
		}

		c := db.sprTopointInPolygonCandidate(ctx, wof_spr)
		candidates = append(candidates, c)
	}

	if err != nil {
		return nil, err
	}

	return candidates, nil
}

func (db *DuckDBSpatialDatabase) PointInPolygonWithChannels(ctx context.Context, spr_ch chan spr.StandardPlacesResult, err_ch chan error, done_ch chan bool, coord *orb.Point, filters ...spatial.Filter) {

	defer func() {
		done_ch <- true
	}()

	for wof_spr, err := range db.pointInPolygon(ctx, coord, filters...) {

		if err != nil {
			err_ch <- err
			break
		}

		spr_ch <- wof_spr
	}
}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidatesWithChannels(ctx context.Context, candidate_ch chan *spatial.PointInPolygonCandidate, err_ch chan error, done_ch chan bool, coord *orb.Point, filters ...spatial.Filter) {

	defer func() {
		done_ch <- true
	}()

	for wof_spr, err := range db.pointInPolygon(ctx, coord, filters...) {

		if err != nil {
			err_ch <- err
			break
		}

		candidate_ch <- db.sprTopointInPolygonCandidate(ctx, wof_spr)
	}

}

func (db *DuckDBSpatialDatabase) Disconnect(ctx context.Context) error {
	return db.conn.Close()
}

func (db *DuckDBSpatialDatabase) sprTopointInPolygonCandidate(ctx context.Context, spr_r spr.StandardPlacesResult) *spatial.PointInPolygonCandidate {

	bounds := orb.Bound{
		Min: [2]float64{spr_r.MinLongitude(), spr_r.MinLatitude()},
		Max: [2]float64{spr_r.MaxLongitude(), spr_r.MaxLatitude()},
	}

	c := &spatial.PointInPolygonCandidate{
		Id:        spr_r.Id(),
		FeatureId: spr_r.Id(),
		IsAlt:     false,
		AltLabel:  "",
		Bounds:    bounds,
	}

	return c
}

func (db *DuckDBSpatialDatabase) pointInPolygon(ctx context.Context, coord *orb.Point, filters ...spatial.Filter) iter.Seq2[spr.StandardPlacesResult, error] {

	logger := slog.Default()
	logger = logger.With("coordinate", coord)

	return func(yield func(spr.StandardPlacesResult, error) bool) {

		yield(nil, fmt.Errorf("Not implemented."))

		// See what's happening here? We are violating the cardinal rule of SQL
		// wrangling by NOT passing in variables to be quoted or escaped as necessary
		// by the SQL driver. That's because every time I do I get errors like this:
		// 2024/12/12 10:46:44 ERROR Q error="Binder Error: Referenced column \"POINT(? ?)\" not found in FROM clause!\nCandidate bindings: \"read_parquet.id\"\nLINE 1: ... WHERE ST_Contains(geometry::GEOMETRY, \"POINT(? ?)\"::GEOMETRY)\n

		q := fmt.Sprintf(`SELECT id, parent_id, name, placetype, country, repo, lat, lon, min_lat, min_lon, max_lat, max_lon, modified FROM read_parquet('%s') WHERE ST_Contains(geometry::GEOMETRY, 'POINT(%f %f)'::GEOMETRY)`, db.database_uri, coord.X(), coord.Y())

		rows, err := db.conn.QueryContext(ctx, q)

		if err != nil {
			logger.Error("Query failed", "error", err)
			yield(nil, err)
			return
		}

		defer rows.Close()

		for rows.Next() {

			var id int64
			var parent_id int64
			var name string
			var placetype string
			var country string
			var repo string
			var lat float64
			var lon float64
			var min_lat float64
			var min_lon float64
			var max_lat float64
			var max_lon float64
			var str_lastmod string

			err := rows.Scan(&id, &parent_id, &name, &placetype, &country, &repo, &lat, &lon, &min_lat, &min_lon, &max_lat, &max_lon, &str_lastmod)

			if err != nil {
				logger.Error("Row scanning failed", "error", err)
				yield(nil, err)
				break
			}

			rel_path, err := uri.Id2RelPath(id)

			if err != nil {
				logger.Error("Failed to derive rel path for ID", "id", id, "error", err)
				yield(nil, err)
				break
			}

			lastmod := int64(0)

			t, err := time.Parse(time.RFC3339, str_lastmod)

			if err != nil {
				logger.Warn("Failed to parse lastmod string", "lastmod", str_lastmod, "error", err)
			} else {
				lastmod = t.Unix()
			}

			wof_spr := &spr.WOFStandardPlacesResult{
				WOFId:           id,
				WOFParentId:     parent_id,
				WOFName:         name,
				WOFPlacetype:    placetype,
				WOFCountry:      country,
				WOFRepo:         repo,
				WOFPath:         rel_path,
				MZURI:           rel_path,
				MZLatitude:      lat,
				MZLongitude:     lon,
				MZMinLatitude:   min_lat,
				MZMinLongitude:  min_lon,
				MZMaxLatitude:   max_lat,
				MZMaxLongitude:  max_lon,
				WOFLastModified: lastmod,
				// Things we don't know
				WOFSupersededBy: make([]int64, 0),
				WOFSupersedes:   make([]int64, 0),
				WOFBelongsTo:    make([]int64, 0),
				MZIsCurrent:     int64(-1),
				MZIsCeased:      int64(-1),
				MZIsDeprecated:  int64(-1),
				MZIsSuperseded:  int64(-1),
				MZIsSuperseding: int64(-1),
			}

			filters_ok := true

			for _, f := range filters {

				err = filter.FilterSPR(f, wof_spr)

				if err != nil {
					logger.Debug("Filter error", "error", err)
					filters_ok = false
					break
				}
			}

			if !filters_ok {
				continue
			}

			yield(wof_spr, nil)
		}

		err = rows.Close()

		if err != nil {
			logger.Error("Failed to close rows", "error", err)
			yield(nil, err)
		}
	}

	//
}

/*

D DESCRIBE SELECT * FROM read_parquet('https://data.geocode.earth/wof/dist/parquet/whosonfirst-data-admin-latest.parquet');
┌───────────────┬────────────────────────────────────────────────────────┬──────┬─────┬─────────┬───────┐
│  column_name  │                      column_type                       │ null │ key │ default │ extra │
├───────────────┼────────────────────────────────────────────────────────┼──────┼─────┼─────────┼───────┤
│ id            │ VARCHAR                                                │ YES  │     │         │       │
│ parent_id     │ INTEGER                                                │ YES  │     │         │       │
│ name          │ VARCHAR                                                │ YES  │     │         │       │
│ placetype     │ VARCHAR                                                │ YES  │     │         │       │
│ placelocal    │ VARCHAR                                                │ YES  │     │         │       │
│ country       │ VARCHAR                                                │ YES  │     │         │       │
│ repo          │ VARCHAR                                                │ YES  │     │         │       │
│ lat           │ DOUBLE                                                 │ YES  │     │         │       │
│ lon           │ DOUBLE                                                 │ YES  │     │         │       │
│ min_lat       │ DOUBLE                                                 │ YES  │     │         │       │
│ min_lon       │ DOUBLE                                                 │ YES  │     │         │       │
│ max_lat       │ DOUBLE                                                 │ YES  │     │         │       │
│ max_lon       │ DOUBLE                                                 │ YES  │     │         │       │
│ min_zoom      │ VARCHAR                                                │ YES  │     │         │       │
│ max_zoom      │ VARCHAR                                                │ YES  │     │         │       │
│ min_label     │ VARCHAR                                                │ YES  │     │         │       │
│ max_label     │ VARCHAR                                                │ YES  │     │         │       │
│ modified      │ DATE                                                   │ YES  │     │         │       │
│ is_funky      │ VARCHAR                                                │ YES  │     │         │       │
│ population    │ VARCHAR                                                │ YES  │     │         │       │
│ country_id    │ VARCHAR                                                │ YES  │     │         │       │
│ region_id     │ VARCHAR                                                │ YES  │     │         │       │
│ county_id     │ VARCHAR                                                │ YES  │     │         │       │
│ concord_ke    │ VARCHAR                                                │ YES  │     │         │       │
│ concord_id    │ VARCHAR                                                │ YES  │     │         │       │
│ iso_code      │ VARCHAR                                                │ YES  │     │         │       │
│ hasc_id       │ VARCHAR                                                │ YES  │     │         │       │
│ gn_id         │ VARCHAR                                                │ YES  │     │         │       │
│ wd_id         │ VARCHAR                                                │ YES  │     │         │       │
│ name_ara      │ VARCHAR                                                │ YES  │     │         │       │
│ name_ben      │ VARCHAR                                                │ YES  │     │         │       │
│ name_deu      │ VARCHAR                                                │ YES  │     │         │       │
│ name_eng      │ VARCHAR                                                │ YES  │     │         │       │
│ name_ell      │ VARCHAR                                                │ YES  │     │         │       │
│ name_fas      │ VARCHAR                                                │ YES  │     │         │       │
│ name_fra      │ VARCHAR                                                │ YES  │     │         │       │
│ name_heb      │ VARCHAR                                                │ YES  │     │         │       │
│ name_hin      │ VARCHAR                                                │ YES  │     │         │       │
│ name_hun      │ VARCHAR                                                │ YES  │     │         │       │
│ name_ind      │ VARCHAR                                                │ YES  │     │         │       │
│ name_ita      │ VARCHAR                                                │ YES  │     │         │       │
│ name_jpn      │ VARCHAR                                                │ YES  │     │         │       │
│ name_kor      │ VARCHAR                                                │ YES  │     │         │       │
│ name_nld      │ VARCHAR                                                │ YES  │     │         │       │
│ name_pol      │ VARCHAR                                                │ YES  │     │         │       │
│ name_por      │ VARCHAR                                                │ YES  │     │         │       │
│ name_rus      │ VARCHAR                                                │ YES  │     │         │       │
│ name_spa      │ VARCHAR                                                │ YES  │     │         │       │
│ name_swe      │ VARCHAR                                                │ YES  │     │         │       │
│ name_tur      │ VARCHAR                                                │ YES  │     │         │       │
│ name_ukr      │ VARCHAR                                                │ YES  │     │         │       │
│ name_urd      │ VARCHAR                                                │ YES  │     │         │       │
│ name_vie      │ VARCHAR                                                │ YES  │     │         │       │
│ name_zho      │ VARCHAR                                                │ YES  │     │         │       │
│ geom_src      │ VARCHAR                                                │ YES  │     │         │       │
│ geometry      │ BLOB                                                   │ YES  │     │         │       │
│ geometry_bbox │ STRUCT(xmin FLOAT, ymin FLOAT, xmax FLOAT, ymax FLOAT) │ YES  │     │         │       │
└───────────────┴────────────────────────────────────────────────────────┴──────┴─────┴─────────┴───────┘

*/
