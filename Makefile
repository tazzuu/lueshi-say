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

# Build and push 'latest' ;
# make docker-build docker-push GIT_TAG=latest
# NOTE: this only builds for the current platform arch, use buildx instead for multi-platform ; see docs listed
DOCKER_TAG:=tazzuu/lueshi-say:$(GIT_TAG)
docker-build:
	docker build --build-arg "Version=$(GIT_TAG)" -t $(DOCKER_TAG) .

docker-push:
	docker push $(DOCKER_TAG)

# make multi-arch Docker build;
# make sure to do ALL the commands listed here in order https://cloudolife.com/2022/03/05/Infrastructure-as-Code-IaC/Container/Docker/Docker-buildx-support-multiple-architectures-images/
# make docker-buildx-push GIT_TAG=latest
docker-buildx-push:
	docker buildx build --platform linux/arm/v7,linux/arm64,linux/amd64 --push --build-arg "Version=$(GIT_TAG)" --tag $(DOCKER_TAG) .