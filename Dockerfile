FROM golang

ARG ARG_PORT
ENV PORT=$ARG_PORT

COPY . /video-downloader

RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl \
 && chmod a+rx /usr/local/bin/youtube-dl

WORKDIR /video-downloader
RUN go build

EXPOSE $PORT

CMD ./video-downloader $PORT