FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /gobank

EXPOSE 3000

CMD [ "/gobank" ]