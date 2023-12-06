#!/bin/sh

ps -aux|grep 'go-build'| grep 'exe/server' | grep -v grep| awk '{print $2}'|xargs kill -9 2>/dev/null
cd /root/proxy-autoport
git pull origin master
chmod a+x /root/proxy-autoport/resource/server/*
nohup go run main/server.go &