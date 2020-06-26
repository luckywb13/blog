FROM alpine:3.9

RUN mkdir -p /data/conf

COPY main /data/

COPY views /data/views

COPY static /data/static

COPY app.yaml /data/conf/app.yaml

COPY wait-for-it.sh /data/wait-for-it.sh

RUN chmod +x /data/wait-for-it.sh

RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

WORKDIR /data/