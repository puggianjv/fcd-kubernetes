FROM golang:1.15.2-alpine as builder
WORKDIR /go/src/app
ENV CGO_ENABLED=0
COPY . .
RUN go build -o main .

FROM scratch
COPY --from=builder /go/src/app/ /
ENTRYPOINT ["/main"]