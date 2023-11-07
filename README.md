# Golang Youtube Stream

Golang audio streaming server from youtube videos only using core packages of golang.

---

## Run locally


Build docker image

```bash

docker build -t golang-youtube-stream .

```

Run docker image

```bash

docker run -p 8080:8080 golang-youtube-stream

```

---


## Usage

Example youtube video: https://www.youtube.com/watch?v=yLOM8R6lbzg.

You can listen audio stream from youtube video using the following url:

```bash

http://localhost:8080/watch?v=yLOM8R6lbzg

```







