FROM golang:1.14-alpine as builder

WORKDIR /go-yt

COPY . .


RUN go build -o /main .

FROM alpine

RUN apk add --no-cache youtube-dl ffmpeg

COPY --from=builder /main /main


ENTRYPOINT [ "/main" ]





