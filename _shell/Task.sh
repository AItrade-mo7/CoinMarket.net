#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

StartName="Task-TickerAnaly"
BuildName=${StartName}
DeployPath="/root/ProdProject/"${StartName}

echo "开始打包" &&
  npm run build

echo "停止 pm2 服务" &&
  pm2 delete "${StartName}"

echo "移动文件到 ProdProject 目录"
cp -r "${OutPutPath}/." "${DeployPath}/"

cd "${DeployPath}" || exit

echo "启动 pm2 服务"
pm2 start "./${BuildName}" --name "${StartName}" &&
  exit 0
