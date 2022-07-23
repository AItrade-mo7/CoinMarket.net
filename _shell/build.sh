#!/bin/bash
source "./_shell/init.sh"
#############

echo " =========== 正在进行编译 =========== "

set GOARCH=amd64
go mod tidy &&
  go build -o ${buildName}
echo "编译 完成"

echo " =========== 开始进行 文件整理 =========== "

echo "清理并创建 dist 目录"
rm -rf ${outPutPath}
mkdir ${outPutPath} &&
  echo "移动 goRun 文件"
mv ${buildName} ${outPutPath}
