# linux 64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./release/beekey-linux64
# windows 64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./release/beekey-win64
# mac 64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./release/beekey-mac64