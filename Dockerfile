FROM golang as builder

# Path in builders $GOPATH (/go)
WORKDIR /go/src/github.com/didip/shawty
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /go/src/github.com/didip/shawty/main /app/
WORKDIR /app
CMD ["./main"]
