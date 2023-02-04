#!/bin/bash
source "./_shell/init.sh"
#############

echo " =========== go build  =========== "

go mod tidy &&
  go build -o "${BuildName}"

echo " server 端编译 完成 "

echo " =========== 开始进行文件整合 =========== "

rm -rf "${OutPutPath}"
mkdir "${OutPutPath}"

echo "移动 go build 文件"
mv "${BuildName}" "${OutPutPath}/" &&
  cp -r "${NowPath}/ReStart.sh" "${OutPutPath}/"
exit
