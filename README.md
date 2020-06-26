blog
==============
`luckywb13/blog` gin写的个人博客,db为mysql(v>=5.6),前端小白,将就能用

---------------------------------------

## 地址
https://www.luckywb.com

## 编译
```sh
$ git clone https://github.com/luckywb13/blog.git
$ cd blog
$ make linux_build
```

## Docker启动
```sh
$ docker-compose up -d --build
```

## 手动启动
```sh
$ ./main -debug=false -conf=app.yaml
```
手动启动时需要手动创建数据表,表结构文件为blog.sql,然后修改app.yaml中db的相关配置。

## 访问方式
后端管理页面:http://localhost:8080/admin/login

博客主页面:http://localhost:8080

后台初始账号:admin@qq.com 密码:123456

留言版引用畅言平台, 需要注册并且审核备案号后获取信息
