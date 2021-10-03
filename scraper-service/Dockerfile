FROM golang:alpine as builder
WORKDIR /go/src/github.com/Ayush-Walia/amazon-scraper/scraper-service
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cg -o /scraper-service main.go

FROM scratch
COPY --from=builder /scraper-service scraper-service
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080