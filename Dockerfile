FROM golang:1.16.5-alpine

WORKDIR $GOPATH/src/github.com/b-nova-openhub/fisicus/fisicus

COPY . .

RUN go get -d -v ./... \
    && go build -o bin/fisicus main.go \
    && go install -v ./... \

EXPOSE 8080

CMD ["sh", "fisicus"]