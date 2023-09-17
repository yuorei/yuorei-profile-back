FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN go build  -o blog
CMD [ "./blog" ]

EXPOSE 8080
