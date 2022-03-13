#!/usr/bin/bash
read -p "请输入注释:" input
if [$input == null] || [$input == ""];then
        echo "您的输入为$input ,系统将使用默认注释"
        git add .&&git commit -m "自动更新"&&git push origin master
else
        git add .&&git commit -m "$input"&&git push origin master
fi
echo "上传完成"