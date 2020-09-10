FROM golang

ARG ARG_PORT
ENV PORT=$ARG_PORT

COPY . /video-downloader
ADD ./resources/user-agent.txt /root/.cache/instalooter/2.4.4/user-agent.txt

# Installing Instagram tool
RUN apt update -y \
 && apt upgrade -y \
 && apt install -y python3 python3-pip \
 && pip3 install --user instalooter --pre

# Installing Youtube tool
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl \
 && chmod a+rx /usr/local/bin/youtube-dl

ENV PATH="/root/.local/bin:${PATH}"

WORKDIR /video-downloader
RUN go build

EXPOSE $PORT

CMD ./video-downloader $PORT