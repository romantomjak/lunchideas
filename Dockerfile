FROM golang

WORKDIR /opt/lunchideas

COPY foursquare.go main.go go.mod go.sum ./

RUN go mod download

RUN go install

CMD /go/bin/lunchideas
