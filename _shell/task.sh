#!/bin/bash
# 加载变量
source "./_shell/init.sh"
#############

startName="TickerAnaly.Task"
deployPath="/root/ProdProject/TickerAnaly"

echo "开始打包" &&
  npm run build

echo "停止 pm2 服务" &&
  pm2 delete ${startName}

rm -rf ${deployPath}
echo "移动文件到 ProdProject 目录"
cp -r ${outPutPath}"/." ${deployPath}"/"

cd ${deployPath}

echo "启动 pm2 服务"
pm2 start ./${buildName} --name ${startName} --no-autorestart
