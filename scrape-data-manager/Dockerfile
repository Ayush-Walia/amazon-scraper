FROM golang:alpine as builder
WORKDIR /go/src/github.com/Ayush-Walia/amazon-scraper/scrape-data-manager
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cg -o /scrape-data-manager main.go

FROM scratch
COPY --from=builder /scrape-data-manager scrape-data-manager
EXPOSE 8080