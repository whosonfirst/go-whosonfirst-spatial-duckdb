package duckdb

import (
	"context"
	"database/sql"
	"fmt"
	"iter"
	"net/url"

	"github.com/paulmach/orb"
	"github.com/whosonfirst/go-whosonfirst-spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type DuckDBSpatialDatabase struct {
	database.SpatialDatabase
	conn         *sql.DB
	database_uri string
}

func NewDuckDBSpatialDatabase(ctx context.Context, uri string) (database.SpatialDatabase, error) {

	_, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	conn, err := sql.Open("duckdb", "")

	if err != nil {
		return nil, fmt.Errorf("Failed to create database connection, %w", err)
	}

	extensions := []string{
		"spatial",
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
		database_uri: uri,
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

	for spr, err := range db.pointInPolygon(ctx, coord, filters...) {

		if err != nil {
			return nil, err
		}

		results = append(results, spr)
	}

	return nil, fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidates(ctx context.Context, centroid *orb.Point, filters ...spatial.Filter) ([]*spatial.PointInPolygonCandidate, error) {

	return nil, fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) PointInPolygonWithChannels(ctx context.Context, spr_ch chan spr.StandardPlacesResult, err_ch chan error, done_ch chan bool, coord *orb.Point, filters ...spatial.Filter) {

	defer func() {
		done_ch <- true
	}()

	for spr, err := range db.pointInPolygon(ctx, coord, filters...) {

		if err != nil {
			err_ch <- err
			break
		}

		spr_ch <- spr
	}
}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidatesWithChannels(ctx context.Context, candidate_ch chan *spatial.PointInPolygonCandidate, err_ch chan error, done_ch chan bool, pt *orb.Point, filters ...spatial.Filter) {

}

func (db *DuckDBSpatialDatabase) Disconnect(ctx context.Context) error {
	return db.conn.Close()
}

func (db *DuckDBSpatialDatabase) pointInPolygon(ctx context.Context, coord *orb.Point, filters ...spatial.Filter) iter.Seq2[spr.StandardPlacesResult, error] {

	return func(yield func(spr.StandardPlacesResult, error) bool) {

		yield(nil, fmt.Errorf("Not implemented."))

		q := "SELECT id FROM read_parquet('?') WHERE ST_Contains(geometry::GEOMETRY, 'POINT(? ?)'::GEOMETRY)"

		rows, err := db.conn.QueryContext(ctx, q, coord.X, coord.Y)

		if err != nil {
			yield(nil, err)
			return
		}

		defer rows.Close()

		for rows.Next() {

			var id int64

			err := rows.Scan(&id)

			if err != nil {
				yield(nil, err)
				break
			}

			// Derive SPR here...
		}

		err = rows.Close()

		if err != nil {
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
