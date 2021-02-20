FROM golang:alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir /build
WORKDIR /build
COPY /application/*.go .
RUN go build -a -ldflags '-w -extldflags "-static"' -o server *.go

FROM scratch

EXPOSE 8000

COPY --from=builder /build/server .
COPY /application/count.gob .
COPY /application/www ./www

CMD ["./server"]
