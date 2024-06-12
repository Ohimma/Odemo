#!/bin/bash

# git remote add github git@github.com:Ohimma/Odemo.git
# git remote add gitee  git@gitee.com:ohimma/Odemo.git

mes=$1

if [[ -z $1 ]];then
   echo "请输入commit提交信息...."
   exit
else
   echo "\033[1;34mgit commit -m $mes \033[0m"
   git add --all .
   git commit -m "$mes"
   git push github main
   git push gitee  main
fi
