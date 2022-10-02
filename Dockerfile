FROM golang:1.19 AS build

WORKDIR /app

COPY . /app

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o api cmd/main.go

FROM scratch 

WORKDIR /app

COPY --from=build /app/api ./

EXPOSE 9000

CMD ["./api"]

