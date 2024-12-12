# https://github.com/marcboeker/go-duckdb?tab=readme-ov-file#vendoring
# go install github.com/goware/modvendor@latest
modvendor:
	modvendor -copy="**/*.a **/*.h" -v
