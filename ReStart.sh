#!/bin/bash

# 将数据同步至硬盘
sync

# 清除刷新swap
swapoff -a && swapon -a

# 重启 mongodb
sudo systemctl stop mongod
sudo systemctl restart mongod
