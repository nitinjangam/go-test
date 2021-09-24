FROM golang:1.17-alpine AS builder

WORKDIR /app

RUN apk update && apk add --no-cache git ca-certificates

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /main

FROM gcr.io/distroless/base-debian10 AS runner

WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /main /main

USER 1000:1000

EXPOSE 8080

ENTRYPOINT ["/main"]