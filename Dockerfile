FROM golang

ADD . /go/src/github.com/romantomjak/lunchideas

RUN go install github.com/romantomjak/lunchideas

CMD /go/bin/lunchideas
