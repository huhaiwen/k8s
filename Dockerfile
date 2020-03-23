FROM golang:1.8.0:alpine
ADD . /golang/src/app
WORKDIR /golang/src/app
RUN GOOS=linux GOARCH=386 go build /golang/src/app/main.go
CMD ["./main"]