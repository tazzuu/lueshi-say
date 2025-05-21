# BUILD

FROM --platform=linux/amd64 golang:1.24.3-alpine AS build
COPY ./main.go /main.go
RUN go build -o /lueshi-say /main.go

# DEPLOY

FROM scratch
COPY --from=build /lueshi-say /usr/local/bin/lueshi-say
ENTRYPOINT ["/usr/local/bin/lueshi-say"]