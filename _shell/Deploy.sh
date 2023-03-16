#!/bin/bash

## 设置并加载变量
source "./_shell/init.sh"
BuildName=${BuildName}
StartName=${StartName}
OutPutPath=${OutPutPath}
DeployPath=${DeployPath}

echo "开始打包" &&
  npm run build

echo "停止 pm2 服务" &&
  pm2 delete "${StartName}"

rm -rf "${DeployPath}"
mkdir "${DeployPath}"
echo "移动文件到 ProdProject 目录"
cp -r "${OutPutPath}/." "${DeployPath}/"

cd "${DeployPath}" || exit

echo "启动 pm2 服务"
pm2 start "./${BuildName}" --name "${StartName}" &&
  exit 0
