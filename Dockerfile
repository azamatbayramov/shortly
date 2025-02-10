FROM golang:1.23.6-alpine3.21 AS builder

LABEL author="azamatbayramov"

ENV CGO_ENABLED=0

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /build/app cmd/server/main.go

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /build/app /app/app
COPY html/ /app/html

CMD ["./app"]
