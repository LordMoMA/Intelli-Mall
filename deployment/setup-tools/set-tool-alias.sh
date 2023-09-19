#!/usr/bin/env sh

root=$(realpath ${PWD}/../..)

docker build -t deploytools:latest .

alias deploytools='docker run --rm -it -v /var/run/docker.sock:/var/run/docker.sock -v ~/.aws:/root/.aws -v ~/.docker:/root/.docker -v ~/.kube:/root/.kube -v ${root}:/intellimall -v ${PWD}:/intellimall/deployment/.current -w /intellimall/deployment/.current deploytools'

echo "---"
echo
echo "Usage: deploytools <cmd [parameters]>"
