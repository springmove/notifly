FROM alpine:3.6

RUN apk update && apk add curl
RUN mkdir -p /etc/notifly/conf /etc/notifly/log

COPY ./entrypoint.sh /
COPY ./conf/config.yml /etc/notifly/config.yml
#COPY ./etc/notifly/api.json /etc/notifly
COPY ./build/notifly /usr/bin

ENTRYPOINT ["/entrypoint.sh"]
CMD notifly --config /etc/notifly/conf/config.yml
