#!/bin/sh

echo "执行的文件名：$0";
echo "第一个参数为：$1";

sed -i "43s/^.*$/      \"port\": $1,/" /Users/ld/my-file/workspace/go_own/proxy-autoport/resource/config.json
#monit restart xray
