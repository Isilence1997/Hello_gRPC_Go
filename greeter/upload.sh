#!/bin/bash
#支持交叉编译
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

#以下配置信息需要修改为你当前服务的信息
app=video_app_short_video
server=hello_alice
env=59d03cd1 //服务所在的环境名,http://123.woa.com/59d03cd1#
instances=cls-3qqm0f6l-7ab6779ab56881d63cb3d426960d1b5b-0 // 发布实例名(节点名)
user=alicehyhe //发布人rtx

# 删除旧的文件
rm ${server}
#编译构建生成二进制文件(不需修改)
go build -o ${server}

#发布命令(不需修改) ps:确保你的dtools是安装在/usr/bin/下
dtools bpatch -env "${env}" -app "${app}" -server "${server}" -bin "${server}"  -instances "${instances}" -user "${user}" -lang=go