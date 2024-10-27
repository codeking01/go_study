#!/bin/bash

# 定义一个函数来执行巡检命令
check_command() {
    echo "正在执行命令: \$1"
    eval "\$1"  # 执行传入的命令
    if [ $? -eq 0 ]; then
        echo "命令执行成功."
    else
        echo "命令执行失败."
    fi
}

# 执行巡检命令
check_command "echo $PACKAGE_ROOT"


#sleep 5
cat ttt
xxxx | exit 1
echo "hello exit"
sleep 100
# 如果所有命令都成功，返回 0
exit 0
