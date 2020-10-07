FROM golang:alpine as builder
WORKDIR /go/src/
COPY . .
RUN go build main.go

FROM scratch

COPY --from=builder /go/src/ .

ENTRYPOINT [ "./main" ]