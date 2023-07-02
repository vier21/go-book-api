FROM golang:1.19 AS build

WORKDIR /app
COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o api cmd/app/main.go

FROM build AS dev-env
CMD ["go", "run", "cmd/app/main.go"]

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/api .

ENTRYPOINT [ "./api" ]

