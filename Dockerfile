FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod tidy

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./

EXPOSE 8080

CMD [ "/sirekap-backend" ]