FROM golang as builder

WORKDIR /go/src/github.com/rluisr/shawty
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch
COPY --from=builder /go/src/github.com/rluisr/shawty/main /app/
WORKDIR /app
CMD ["./main"]
