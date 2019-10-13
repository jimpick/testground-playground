#! /bin/bash

docker ps -a | awk '{print $1}' | grep -v CONTAINER | xargs docker rm

docker images | grep -v 'REPOSITORY\|golang\|redis\|busybox' | awk '{print $3}' | xargs docker rmi

for f in `ifconfig | grep veth | awk '{print $1}' | sed 's,:,,'`; do echo $f; sudo ifconfig $f down; done

for f in `ifconfig | grep br- | awk '{print $1}' | sed 's,:,,'`; do echo $f; sudo ifconfig $f down; done
