docker run --rm -v /Users/sqeven/go/src/github.com/sqeven/robot:/go/src/github.com/sqeven/robot -w /go/src/github.com/sqeven/robot golang:1.12.7  go build -o bootstrap example/alifaas/faas.go

mv ../../bootstrap .
zip code.zip bootstrap
fun deploy