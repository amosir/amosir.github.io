---
title: "Mac更换homebrew下载源"
date: 2020-04-26T16:43:38+08:00
draft: false
image: "img/g6Wd6MS8lms.jpg"
categories: 
  - 工具
tag:
---


将下载源切换为清华

```bash
echo 'export HOMEBREW_BOTTLE_DOMAIN=https://mirrors.tuna.tsinghua.edu.cn/homebrew-bottles'>>   ~/.bash_profile
source ~/.bash_profile   #执行.bash_profile脚本让配置即时生效
```

