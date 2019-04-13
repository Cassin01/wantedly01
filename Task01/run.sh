#!/bin/zsh
docker build -t echo .
#docker run --rm -p 127.0.0.1:1323:1323 --name echo echo
docker run --rm -p 127.0.0.1:8080:8080 --name echo echo
