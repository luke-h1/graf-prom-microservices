FROM golang:1.22-alpine

WORKDIR /build

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s -w -extldflags "-static"' -o app .

FROM alpine:3.19.1
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=0 /build/app .

CMD ["./app"]