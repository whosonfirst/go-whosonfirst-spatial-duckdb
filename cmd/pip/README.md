# pip

Perform an point-in-polygon operation for an input latitude, longitude coordinate and on a set of Who's on First records stored in a spatial database.

```
$> ./bin/pip -h
Perform an point-in-polygon operation for an input latitude, longitude coordinate and on a set of Who's on First records stored in a spatial database.
Usage:
	 ./bin/pip [options]
Valid options are:

  -alternate-geometry value
    	One or more alternate geometry labels (wof:alt_label) values to filter results by.
  -cessation string
    	A valid EDTF date string.
  -custom-placetypes string
    	A JSON-encoded string containing custom placetypes defined using the syntax described in the whosonfirst/go-whosonfirst-placetypes repository.
  -enable-custom-placetypes
    	Enable wof:placetype values that are not explicitly defined in the whosonfirst/go-whosonfirst-placetypes repository.
  -geometries string
    	Valid options are: all, alt, default. (default "all")
  -inception string
    	A valid EDTF date string.
  -is-ceased value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-current value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-deprecated value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-superseded value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -is-superseding value
    	One or more existential flags (-1, 0, 1) to filter results by.
  -iterator-uri value
    	Zero or more URIs denoting data sources to use for indexing the spatial database at startup. URIs take the form of {ITERATOR_URI} + "#" + {PIPE-SEPARATED LIST OF ITERATOR SOURCES}. Where {ITERATOR_URI} is expected to be a registered whosonfirst/go-whosonfirst-iterate/v2 iterator (emitter) URI and {ITERATOR SOURCES} are valid input paths for that iterator. Supported whosonfirst/go-whosonfirst-iterate/v2 iterator schemes are: cwd://, directory://, featurecollection://, file://, filelist://, geojsonl://, null://, repo://.
  -latitude float
    	A valid latitude.
  -longitude float
    	A valid longitude.
  -mode string
    	Valid options are: cli, lambda. (default "cli")
  -placetype value
    	One or more place types to filter results by.
  -properties-reader-uri string
    	A valid whosonfirst/go-reader.Reader URI. Available options are: [fs:// null:// repo:// stdin://]. If the value is {spatial-database-uri} then the value of the '-spatial-database-uri' implements the reader.Reader interface and will be used.
  -property value
    	One or more Who's On First properties to append to each result.
  -sort-uri value
    	Zero or more whosonfirst/go-whosonfirst-spr/sort URIs.
  -spatial-database-uri string
    	A valid whosonfirst/go-whosonfirst-spatial/data.SpatialDatabase URI. options are: [duckdb:// rtree://] (default "rtree://")
  -verbose
    	Enable verbose (debug) logging.
```

## Example

```
$> bin/pip \
	-spatial-database-uri 'duckdb://?uri=http://localhost:8080/whosonfirst-data-admin-latest.parquet' \
	-latitude 53.55192911144368 \
	-longitude -113.49058382120987 \
	-placetype locality \
	| jq
	
{
  "results": [
    null,
    {
      "edtf:inception": "",
      "edtf:cessation": "",
      "wof:id": 890458485,
      "wof:parent_id": 1511799791,
      "wof:name": "Edmonton",
      "wof:placetype": "locality",
      "wof:country": "CA",
      "wof:repo": "whosonfirst-data-admin-ca",
      "wof:path": "890/458/485/890458485.geojson",
      "wof:superseded_by": [],
      "wof:supersedes": [],
      "wof:belongsto": [],
      "mz:uri": "890/458/485/890458485.geojson",
      "mz:latitude": 53.546218,
      "mz:longitude": -113.490371,
      "mz:min_latitude": 53.395405,
      "mz:min_longitude": -113.71388,
      "mz:max_latitude": 53.715919,
      "mz:max_longitude": -113.271524,
      "mz:is_current": -1,
      "mz:is_ceased": -1,
      "mz:is_deprecated": -1,
      "mz:is_superseded": -1,
      "mz:is_superseding": -1,
      "wof:lastmodified": 1690848000
    }
  ]
}
```
