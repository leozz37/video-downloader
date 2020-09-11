FROM golang

ARG ARG_PORT
ENV PORT=$ARG_PORT

RUN useradd -ms /bin/bash prod

COPY . /video-downloader
RUN chown -R prod:prod /video-downloader

 # Installing Youtube tool
RUN apt update -y \
 && apt install -y youtube-dl 

WORKDIR /video-downloader
RUN go build

USER prod

EXPOSE $PORT

CMD ./video-downloader $PORT