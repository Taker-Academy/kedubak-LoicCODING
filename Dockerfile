FROM golang:1.22

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build .

EXPOSE 8080

CMD ["./kedubak-LoicCODING"]