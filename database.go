package duckdb

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/whosonfirst/go-whosonfirst-spatial"
	"github.com/whosonfirst/go-whosonfirst-spatial/database"
	"github.com/whosonfirst/go-whosonfirst-spr/v2"
)

type DuckDBSpatialDatabase struct {
	database.SpatialDatabase
	conn *sql.DB
}

func NewDuckDBSpatialDatabase(ctx context.Context, uri string) (database.SpatialDatabase, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (db *DuckDBSpatialDatabase) IndexFeature(context.Context, []byte) error {
	return fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) RemoveFeature(context.Context, string) error {
	return fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) PointInPolygon(ctx context.Context, coord *orb.Point, filters ...spatial.Filter) (spr.StandardPlacesResults, error) {

	return nil, fmt.Errorf("Not implemented.")

	// SELECT id,name FROM read_parquet('/usr/local/whosonfirst-data-admin-latest.parquet') WHERE ST_Contains(geometry::GEOMETRY, 'POINT(-113.49058382120987 53.55192911144368)'::GEOMETRY);
}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidates(ctx context.Context, centroid *orb.Point, filters ...spatial.Filter) ([]*spatial.PointInPolygonCandidate, error) {

	return nil, fmt.Errorf("Not implemented.")
}

func (db *DuckDBSpatialDatabase) PointInPolygonWithChannels(ctx context.Context, spr_ch chan spr.StandardPlacesResult, err_ch chan error, done_ch chan bool, pt *orb.Point, filters ...spatial.Filter) {

}

func (db *DuckDBSpatialDatabase) PointInPolygonCandidatesWithChannels(ctx context.Context, candidate_ch chan *spatial.PointInPolygonCandidate, err_ch chan error, done_ch chan bool, pt *orb.Point, filters ...spatial.Filter) {

}

func (db *DuckDBSpatialDatabase) Disconnect(ctx context.Context) error {
	return db.conn.Close()
}
