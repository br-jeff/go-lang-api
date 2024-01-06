FROM golang:1.20.12

RUN apt update -y && apt install -y libhyperscan-dev

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /usr/local/bin/app
EXPOSE 8030

CMD [ "app" ]