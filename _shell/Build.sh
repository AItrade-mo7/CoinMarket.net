#!/bin/bash

## 设置并加载变量
source "./_shell/init.sh"
BuildName=${BuildName}
OutPutPath=${OutPutPath}

echo " =========== go build  =========== "

go mod tidy &&
  go build -o "${BuildName}"

echo " =========== 开始进行文件整合 =========== "

rm -rf "${OutPutPath}"
mkdir "${OutPutPath}"

echo "移动 go build 文件"
mv "${BuildName}" "${OutPutPath}/" &&
  exit 0
