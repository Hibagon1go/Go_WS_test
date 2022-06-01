FROM golang:1.18

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

#ホットリロードを可能に
CMD ["air"]