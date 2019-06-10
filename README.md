## Golang实现云存储系统

1.基于golang语言实现分布式文件上传服务


2. 结合ceph以及阿里oss支持断点续传及秒传功能


3. 容器部署


## 技术栈

* redis
* rabbitmq
* docker/k8s
* Ceph
* 阿里云oss


## 过程

先实现一个简单的上传下载服务

![](http://ww1.sinaimg.cn/large/006tNc79gy1g3tr9xrj44j30sy0c6dha.jpg)


实现秒传功能

![](https://ws3.sinaimg.cn/large/006tNc79gy1g3w5cv9da4j31ig0rkao1.jpg)

思路:客户端上传文件时需要知道文件的sha1,通过对比sha1,相同的文件就不需要上传,直接返回成功。
