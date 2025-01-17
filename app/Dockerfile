FROM golang:1.22 as build

ENV CGO_ENABLED=0

WORKDIR /app
COPY app/go.mod app/go.sum ./
RUN go mod download
ADD ./app/. .
COPY ./app/*.go ./
RUN go build -o main main.go

FROM golang:1.22
WORKDIR /app
COPY --from=build /app ./
EXPOSE 80
ENV PORT=80
CMD ["/app/main"]
