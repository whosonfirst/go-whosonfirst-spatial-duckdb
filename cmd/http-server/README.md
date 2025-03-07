# http-server

```
$> ./bin/http-server -h
  -authenticator-uri string
    	A valid sfomuseum/go-http-auth URI. (default "null://")
  -cors-allow-credentials
    	Allow HTTP credentials to be included in CORS requests.
  -cors-origin value
    	One or more hosts to allow CORS requests from; may be a comma-separated list.
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-cors
    	Enable CORS headers for data-related and API handlers.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-geojson
    	Enable GeoJSON output for point-in-polygon API calls.
  -enable-gzip
    	Enable gzip-encoding for data-related and API handlers.
  -enable-www
    	Enable the interactive /debug endpoint to query points and display results.
  -initial-view string
    	A comma-separated string indicating the map's initial view. Valid options are: 'LON,LAT', 'LON,LAT,ZOOM' or 'MINX,MINY,MAXX,MAXY'.
  -iterator-uri value
    	Zero or more URIs denoting data sources to use for indexing the spatial database at startup. URIs take the form of {ITERATOR_URI} + "#" + {PIPE-SEPARATED LIST OF ITERATOR SOURCES}. Where {ITERATOR_URI} is expected to be a registered whosonfirst/go-whosonfirst-iterate/v2 iterator (emitter) URI and {ITERATOR SOURCES} are valid input paths for that iterator. Supported whosonfirst/go-whosonfirst-iterate/v2 iterator schemes are: cwd://, directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://.
  -leaflet_point_style string
    	A custom Leaflet style definition for points. This may either be a JSON-encoded string or a path on disk.
  -leaflet_style string
    	A custom Leaflet style definition for geometries. This may either be a JSON-encoded string or a path on disk.
  -log-timings
    	Emit timing metrics to the application's logger
  -map-provider string
    	Valid options are: leaflet, protomaps (default "leaflet")
  -map-tile-uri string
    	A valid Leaflet tile layer URI. See documentation for special-case (interpolated tile) URIs. (default "https://tile.openstreetmap.org/{z}/{x}/{y}.png")
  -path-api string
    	The root URL for all API handlers (default "/api")
  -path-ping string
    	The URL for the ping (health check) handler (default "/health/ping")
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// repo:// stdin://]. If the value is {spatial-database-uri} then the value of the '-spatial-database-uri' implements the reader.Reader interface and will be used.
  -protomaps-theme string
    	A valid Protomaps theme label. (default "white")
  -server-uri string
    	A valid aaronland/go-http-server URI. (default "http://localhost:8080")
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [duckdb:// rtree://] (default "rtree://")
```

## Example

```
$> ./bin/http-server \
		-enable-www \
		-initial-view '-122.384292,37.621131,13' \
		-server-uri http://localhost:8080 \
		-spatial-database-uri 'duckdb://?uri=/usr/local/whosonfirst/go-whosonfirst-spatial-duckdb/fixtures/sf_county.parquet'
		
2025/03/07 10:09:50 INFO Listening for requests address=http://localhost:8080
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-spatial-www