#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

echo "更新依赖"
go mod tidy
echo " ========== 开始运行 goServer ========== "
go run main.go
