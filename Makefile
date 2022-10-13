export CGO_ENABLED=0
build:
	go build -ldflags "-s -w" -o afk .
	upx --best --lzma afk