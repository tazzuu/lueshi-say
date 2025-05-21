# BUILD

# make sure base image multi-arch matches desired buildx multi-arch targets
FROM golang:1.24.3-alpine AS build

# overwrite this at build time with docker arg `docker build --build-arg "Version=12345" ...`
ARG Version=latest

COPY ./main.go /main.go
RUN go build -ldflags="-X 'main.Version=$Version'" -o /lueshi-say /main.go

# DEPLOY

FROM scratch
COPY --from=build /lueshi-say /usr/local/bin/lueshi-say
ENTRYPOINT ["/usr/local/bin/lueshi-say"]