#!/bin/sh
set -e

if [ ! -f "/etc/notifly/conf/config.yml" ];then
    cp /etc/notifly/config.yml /etc/notifly/conf/config.yml
fi

exec "$@"
