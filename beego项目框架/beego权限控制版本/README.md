set GOARCH=amd64
set GOOS=linux
go build


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
