FROM golang:latest 

RUN mkdir /go/src/app

RUN go get -u github.com/golang/dep/cmd/dep

ADD ./main.go /go/src/app
COPY ./Gopkg.toml /go/src/app

WORKDIR /go/src/app 

RUN dep ensure

CMD ["go", "run", "main.go"]