FROM golang:1.22.3

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /favicon_service 

CMD ["/favicon_service"]
