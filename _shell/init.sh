#!/bin/bash

## 存储变量

# 项目根目录
path=$(pwd)

# 项目的名字和编译时的名字
startName=${path##*/}
buildName="goRun-"${startName}

# 静态 www 的目录
wwwPath=${path}"/www"

# log 目录
logPath=${path}"/logs"

# 最终的输出目录
outPutPath=${path}"/dist"

# 配置文件
userEnv=${path}"/user_config.yaml"

# 服务器的 file 目录
fileMo7Path="/root/file.mo7.cc"

# 部署目录
deployPath="/root/ProdProject/"${startName}
