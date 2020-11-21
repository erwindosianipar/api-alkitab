FROM golang:alpine
WORKDIR /work
ADD . .
RUN go test passage/handler/passage_test.go -v
RUN go build -o /bin/api-alkitab .
WORKDIR /
RUN rm -r /work
CMD ["/bin/api-alkitab"]
