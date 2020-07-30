clean:
	rm ./cql_*.go
	rm ./easyjson_*.go
	rm ./jsoniter_*.go
	go mod tidy

cql: ./cql/gen.go
	go generate ./cql
	go mod tidy

easyjson: ./easyjson/gen.go
	go generate ./easyjson
	go mod tidy

jsoniter: ./jsoniter/gen.go
	go generate ./jsoniter
	go mod tidy