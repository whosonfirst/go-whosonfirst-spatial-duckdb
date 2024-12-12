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
