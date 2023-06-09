FROM golang

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o crm_service

CMD ["./Mini Project"]