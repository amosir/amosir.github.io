---
title: "Golang热加载库使用技法"
date: 2023-05-18T09:48:20+08:00
draft: false
image: "luca-bravo-alS7ewQ41M8-unsplash.jpg"
categories: 
  - golang
  - 编程技巧
tag:
---

# 使用air库进行go项目的热加载

##  1.背景

最近基于司内trpc-go框架开发一个业务网关，涉及到路由和鉴权的部分，使用配置文件的方式实现。

由于测试过程中需要频繁改动配置文件(配置新路由、上游服务等)，而每次改动配置都需要重启网关服务(配置文件是在启动时加载的)，这个过程很慢，严重影响了开发效率。

因此需要一种热加载能力，来规避开发过程中的繁琐的编译操作，提高效率。

经过一番调研，决定使用air。

## 2.简介

air是一款基于golang开发的实时热加载工具，通过使用该工具，使得开发人员能专注于coding，而不会被编译过程打断。

项目地址: [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)

截止到此前(2023年5月18日)，air在github上的star数已经达到了10.3K，可见已经得到了很多同学的认可，赶紧学起来吧。

## 3.特性

* 彩色日志输出
* 自定义构建或二进制命令
* 支持忽略子目录
* 支持监听新目录
* 更好的构建过程

## 4.安装

在golang1.18以上的版本中，可以使用`go install`命令进行安装

```shell
go install github.com/cosmtrek/air@latest
```

当然也可以通过脚本或docker安装，详细方式可参考github介绍。

## 5.使用方法

### 5.1 配置初始化

首次使用时，进入项目根目录，执行`air init`命令，会生成配置文件`.air.toml`

### 5.2 配置修改

air会基于上述配置文件进行监听和编译等，我们需要针对项目特点进行配置文件修改。

以下是本业务网关项目的air配置

```toml
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "main --conf conf/trpc_local.yaml"
  cmd = "go build -o ./main ."
  delay = 3
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = ["docs"]
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "yaml"]
  include_file = []
  kill_delay = "0s"
  log = "build-errors.log"
  poll = false
  poll_interval = 0
  rerun = false
  rerun_delay = 500
  send_interrupt = false
  stop_on_error = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  main_only = false
  time = false

[misc]
  clean_on_exit = false

[screen]
  clear_on_rebuild = false
  keep_scroll = true
```

主要改动下面几个配置项:

* include_ext
* exclude_dir
* cmd
* bin

### 5.3 启动

直接在项目根目录下执行`air`即可

### 5.4 效果

```shell
pkg/delivery/trpc_delivery.go has changed
building...
running...
```





