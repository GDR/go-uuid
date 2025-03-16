# Start of Selection
FROM golang:1.24.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o uuidserver .

CMD ["./uuidserver"]
# End of Selection
