#!/bin/bash

mes=$1

if [ -z $1 ];then
   echo "请输入commit提交信息...."
   exit
else
   echo "git commit -m \"$mes\""
   git add --all .
   git commit -m "$mes"
   git push
fi
