#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

echo " =========== 整理文件到 www  =========== "
rm -rf ${staticPath}
mv ${outPutPath} ${staticPath}

echo " =========== 写入文件 =========== "
sudo cat >${staticPath}"/index.go" <<END
package www

import "embed"

//go:embed *
var Static embed.FS
END

echo " =========== go build  =========== "

go mod tidy &&
  go build -o ${buildName}
echo " server 端编译 完成"

echo " =========== 开始进行文件整合 =========== "

mkdir ${outPutPath}

echo "移动 go build 文件"
mv ${buildName} ${outPutPath}"/" &&
  exit
