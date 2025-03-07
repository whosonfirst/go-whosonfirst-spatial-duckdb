GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")
LDFLAGS=-s -w

CWD=$(shell pwd)

DATABASE=duckdb://?uri=$(CWD)/fixtures/sf_county.parquet
INITIAL_VIEW=-122.384292,37.621131,13

# https://github.com/marcboeker/go-duckdb?tab=readme-ov-file#vendoring
# go install github.com/goware/modvendor@latest
modvendor:
	modvendor -copy="**/*.a **/*.h" -v

http-server:
	go run -mod $(GOMOD) cmd/http-server/main.go \
		-enable-www \
		-initial-view '$(INITIAL_VIEW)' \
		-server-uri http://localhost:8080 \
		-spatial-database-uri '$(DATABASE)'
