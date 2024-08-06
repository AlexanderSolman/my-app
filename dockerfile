FROM golang:1.21.6-alpine
COPY /main /main
ENTRYPOINT ["/main"]
