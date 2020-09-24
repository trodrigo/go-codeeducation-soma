FROM golang:1.8-alpine

WORKDIR /go/src/soma
ADD ./src/soma /go/src/soma

RUN go install .

#CMD ["go", "test", "-v"]
#CMD ["go", "run", "add.go"]
CMD ["soma"]
