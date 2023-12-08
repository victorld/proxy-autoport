#!/bin/sh

echo "执行的文件名：$0";
echo "第一个参数为：$1";
echo "第二个参数为：$2";

sed -i "" "s/port:.*type/port: $1, type/g" $2
