# smaller image the better
FROM golang:1.26.6-alpine

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
RUN go build -o /gs-1

EXPOSE 9999

# Run
CMD ["/gs-1"]