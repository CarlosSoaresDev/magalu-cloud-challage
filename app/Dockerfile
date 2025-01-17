FROM golang:1.22 as build

ENV CGO_ENABLED=0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
ADD . .
COPY *.go ./
RUN go build -o main main.go

FROM golang:1.22
WORKDIR /app
COPY --from=build /app ./
EXPOSE 8080
ENV PORT=8080
CMD ["/app/main"]
