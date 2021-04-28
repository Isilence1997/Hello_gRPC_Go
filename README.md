# tRPC-Go HelloWorld协议实现与发布


## 项目介绍README
创建helloworld协议，实现echo服务，在123平台发布，用rick管理pb接口。

## 快速上手 Getting Started
使用者如何快速上手使用本组织/项目

* 创建自己的git仓库
```
vepc new video_app_short_video/hello_alice --desc=备注
```
之后所有有关该服务的开发和代码修改都在该仓库上进行，比如简单的echo服务
* clone到本地，如：
```
git clone git@git.code.oa.com:video_app_short_video/hello_alice.git
```
* 初始化 go mod 文件：
```
1. cd hello_alice  # 之后所有的操作都在这个目录下面执行
2. go mod init git.code.oa.com/video_app_short_video/hello_alice
```
### 创建蓝盾流水线
vepc new就会建立git与蓝盾项目的映射，设置流水线回调，对于手动创建的git，可以通过`vepc check`和`vepc fix`配置回调。

#### pb接口
* 我们用protobuf定义服务方法，请求参数和响应参数。可以在rick平台新建PB。选择应用名为video_app_short_video，服务名为hello_alice。PB名称最好与服务名一致，不一致option会自动拼接hello_alice_greeter。应用名和服务名建议都使用小写字母
* 最好将service-name对应的“ReplaceMe”修改为与sever一致（我的不一致）

#### 创建蓝盾流水线
```
vepc create --rick-id=20547 --git-path=video_app_short_video/hello_alice 
```
可以手动选择流水线设置，不要勾选单元测试
### 创建服务
进入123平台测试环境，创建服务，选择所属业务，应用名为 video_app_short_video，服务名为hello_alice。(必须和vepc的应用名和服务名配置的一致，否则蓝盾流水线在测试环境找不到对应服务)
* 代码生成，选择TRPC-Go Stub Mod只生成桩代码，并在123平台上进行go mod。选择“TRPC-Go服务生成” 生成trpc-go代码模板。
* 把tRPC接口->TRPC-GO-服务生成->查看详细使用方法->点我下载，下载下来的文件，push到之前git clone的仓库中。(运行vepc create时)
    * 删除本地的stub文件夹，将go.mod中以下两行删除
    ```
    replace git.code.oa.com/... => ./stub/git.code.oa.com/...
    git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter v0.0.0-0001010100000-0000000
    ```
    * goland终端运行go get -v，或手动粘贴`git.code.oa.com/trpcprotocol/video_app_short_video/hello_alice_greeter v1.1.2`（远端stub文件）

### 开发
在本地仓库修改代码，提交到git中，并更新tRPC接口的代码生成，重新构建镜像，发布节点。
* 北极星的名字服务信息即为接口测试中的target，发送请求到该targer即可触发trpc服务端对端的响应。点击查看IP和端口号

## 常见问题 FAQ
本组织/项目的常见通用问题和官方解答
## 行为准则 Code Of Conduct
本README遵循开源治理指标体系介绍-文档质量要求。
## 如何加入 How To Join
本组织/项目有明确的如何加入和贡献的文字说明
## 团队介绍 Members
本组织/项目的角色分工、人名和联络方式、官方交流/沟通渠道
* 邮件(alicehyhe@tencent.com)

## 参考References 
* [README.md规范](https://iwiki.woa.com/pages/viewpage.action?pageId=289511054)
* [视频后台研发手册](https://git.code.oa.com/videobase/videonavim)
* [trpc-go服务发布123平台和触发蓝盾流水线流程](https://iwiki.woa.com/pages/viewpage.action?pageId=551449221)
