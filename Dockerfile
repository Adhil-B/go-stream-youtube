FROM golang:1.20-alpine as builder

WORKDIR /go-yt

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM alpine

RUN apk add --no-cache youtube-dl ffmpeg

COPY --from=builder /main /main


ENTRYPOINT [ "/main" ]





