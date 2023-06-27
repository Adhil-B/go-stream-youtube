FROM golang

WORKDIR /go-yt

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8080

CMD ["sudo", "apt-get","install", "ffmpeg";"sudo", "apt-get", "install", "youtube-dl";"./main"]



