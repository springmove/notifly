#!/bin/bash

version="0.0.8"

docker_repo="linshenqi/notifly"

docker build -t ${docker_repo}:${version} -t ${docker_repo}:latest .

docker push ${docker_repo}:${version}
docker push ${docker_repo}:latest
