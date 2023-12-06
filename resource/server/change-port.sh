#!/bin/sh

echo "执行的文件名：$0";
echo "第一个参数为：$1";
echo "第一个参数为：$2";

sed -i "43s/^.*$/      \"port\": $1,/" $2
#monit restart xray
