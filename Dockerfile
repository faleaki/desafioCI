FROM golang:1.14
RUN mkdir /go/src/soma
WORKDIR /go/src/soma
COPY . .
RUN GOOS=linux go build  main.go
ENV PATH=$PATH:/bin
#RUN go build  main.go
#ENTRYPOINT ["./main"]
