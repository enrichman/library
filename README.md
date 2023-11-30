go test -c ./cmd/server/. -cover -covermode=set --coverprofile=coverage.out -coverpkg=./...

./server.test -test.coverprofile coverage.out

go tool cover -html=coverage.out
