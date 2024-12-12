# go-whosonfirst-spatial-duckdb

Go package to implement the `whosonfirst/go-whosonfirst-spatial` interfaces using a GeoParquet database.

## Example

```
$> go run cmd/pip/main.go \
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