#!/bin/bash

## 存储变量

# 项目根目录
path=$(pwd)

# 项目的名字和编译时的名字
startName=${path##*/}
buildName=${startName}

# 最终的输出目录
outPutPath=${path}"/dist"

# 前端文件输出目录
staticPath=${path}"/www"

# 部署目录
deployPath="/root/ProdProject/CoinMarket.net"
