FROM golang
WORKDIR /work
ADD . .
RUN passage/handler/passage_test.go -v
RUN go build -o /bin/api-alkitab .
WORKDIR /
RUN rm -r /work
CMD ["/bin/api-alkitab"]