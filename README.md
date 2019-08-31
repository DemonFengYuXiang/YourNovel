# YourNovel - 开源小说搜索引擎

## 起源

  项目起源于Github上另外一个类似的小说搜索引擎项目 https://github.com/howie6879/owllook, 该项目给小说爱好者带来极大的福利(无广告、清爽的界面)。但是, 该项目作者貌似马上要关闭网站了, 想想自己手上有多余的服务器并且刚学完Golang, 马上行动起来使用Golang花了几天的时间开发出了这个项目。
  
  **注意: 本项目前端代码以及服务端的一些思想借鉴了owllook项目,后期可能会根据自己喜好进行修改,因此,若你看见类似的东西请不要惊讶**

## 要求

- Go 1.12+ 
- Go Module 开启
- Redis

## 安装以及使用

- git clone https://github.com/DemonFengYuXiang/YourNovel.git 

- `mv YourNovel yournovel`

- `vim yournovel/yournovel/conf/server.go` 填写Reids地址以及密码等信息

- `cd yournovel`

- ` go run main.go`

## Reference

YourNovel使用了以下第三方包:

- gin
- go-redis/redis
- colly
- goquery 
.... 

YourNovel借鉴项目:

[owllook](https://github.com/howie6879/owllook)

