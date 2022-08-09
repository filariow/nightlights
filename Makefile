.PHONY: build flash

flash:
	tinygo flash -target=arduino main.go

build: *.go *.mod
	tinygo build -target=arduino -o bin/nightlights main.go
