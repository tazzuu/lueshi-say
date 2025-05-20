format:
	gofmt -l -w main.go

SRC:=main.go
BIN:=lueshi-say
build:
	go build -o ./$(BIN) ./$(SRC)
.PHONY:build

# need at least 1 tag for this to work
GIT_TAG:=$(shell git describe --tags)
build-all:
	mkdir -p build ; \
	docker run --rm -ti -v $(PWD):/lueshi-say --workdir /lueshi-say golang:1.24.3 bash -c ' \
	for os in darwin linux windows; do \
	for arch in amd64 arm64; do \
	output="build/$(BIN)-v$(GIT_TAG)-$$os-$$arch" ; \
	if [ "$${os}" == "windows" ]; then output="$${output}.exe"; fi ; \
	echo "building: $$output" ; \
	GOOS=$$os GOARCH=$$arch go build -o "$${output}" $(SRC) ; \
	done ; \
	done \
	'