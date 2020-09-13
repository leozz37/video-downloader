# Go-Video

![Go](https://github.com/leozz37/video-downloader/workflows/Go/badge.svg?branch=master)
![Docker](https://github.com/leozz37/video-downloader/workflows/Docker/badge.svg?branch=master)
![CodeQL](https://github.com/leozz37/video-downloader/workflows/CodeQL/badge.svg?branch=master)

Download videos from YouTube, Twitter, Instagram, Facebook and Twitch!

You can access the [website](https://go-video.herokuapp.com/) and download any video!

Under development. Made for study porpouses.

## Running

To run the backend, cd into api directory and run these commands:

```
$ go mod download

$ go run server.go 8090
```

To run the frontend, cd into web directory and run these commands:

```
$ npm install

$ npm run serve
```

## Running containers

To run the backend, cd into api directory and run these commands:

```
$ docker build . --build-arg ARG_PORT=8090 -t govideo-api:latest  

$ docker run -p 8090:8090 -ti govideo-api:latest
```

To run the frontend, cd into web directory and run these commands:

```
$ docker build . --build-arg ARG_PORT=8080 -t govideo-web:latest 

$ docker run -v ${PWD}:/app -v /app/node_modules -p 8080:8080 --rm govideo-web:latest
```