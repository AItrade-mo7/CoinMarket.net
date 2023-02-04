#!/bin/bash
source "./_shell/init.sh"
#############

echo " =========== go build  =========== "

go mod tidy &&
  go build -o ${buildName}
echo " server 端编译 完成"

echo " =========== 开始进行文件整合 =========== "
rm -rf ${outPutPath}

mkdir ${outPutPath}

echo "移动 go build 文件"
mv ${buildName} ${outPutPath}"/" &&
  cp -r ${path}"/ReStart.sh" ${outPutPath}"/"
exit
