FROM golang

# Getting backend port from env
ARG PORT_API
ENV PORT=$PORT_API

# Creating production user
RUN useradd -ms /bin/bash prod

# Copying files
COPY . /video-downloader

# Giving permission to source directory
RUN chown -R prod:prod /video-downloader

# Installing Youtube tool
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl \
 && chmod a+rx /usr/local/bin/youtube-dl

# Logging as production user
USER prod

# Exposing backend port
EXPOSE $PORT

# Installing Go deps
WORKDIR /video-downloader
RUN go mod download

# Running server
CMD go run server.go $PORT