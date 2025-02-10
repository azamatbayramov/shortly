FROM golang:1.23.6-alpine3.21@sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037 AS builder

LABEL author="azamatbayramov"

ENV CGO_ENABLED=0

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o /build/app cmd/server/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /build/app /app/app

CMD ["./app"]
