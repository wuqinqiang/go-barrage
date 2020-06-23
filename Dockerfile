FROM golang:latest

WORKDIR $GOPATH/src/github.com/wuqinqiang/go-barrage
COPY . $GOPATH/src/github.com/wuqinqiang/go-barrage
RUN go build chitchat

EXPOSE 8000
ENTRYPOINT ["./chitchat"]
