---
title: "Hexo NexT博客系统"
date: 2022-04-28T20:27:52+08:00
draft: false
image: "img/xJ2tjuUHD9M.jpg"
categories: 
  - 工具
tag:
---

#! https://zhuanlan.zhihu.com/p/477004737

> 使用 hexo 搭建博客系统需要机器上已经安装 node 和 git

## 基础安装

### 配置 NPM 下载源

为了加快 npm 下载包的速度可以为 npm 配置国内镜像源，后面使用 npm 命令的地方都一致使用`cnpm`

```shell
npm install -g cnpm --registry=https://registry.npm.taobao.org

#下面这一行检测cnpm是否安装成功
cnpm -v
```

### 安装 Hexo

在终端通过 cnpm 安装 hexo:

```shell
cnpm install -g hexo-cli
#查看hexo是否安装成功
hexo -version
```

#### 运行 hexo

```shell
#初始化
hexo init
```

执行以下命令:

```shell
# 使用调试模式启动,这样有些情况下修改后不用重新启动
hexo s --debug
```

下面就可以通过宿主机浏览器访问 [http://localhost:4000](http://localhost:4000)

#### hexo 常用命令

```shell
hexo g    # 等同于hexo generate
hexo s    # 等同于hexo server
hexo p    # 等同于hexo port 修改端口(默认4000)
hexo d    # 等同于hexo deploy
hexo s -g # 组合使用
```

### 安装 NexT 主题

在博客根目录(blog)下安装 next 主题

```shell
git clone https://github.com/theme-next/hexo-theme-next.git themes/next
```

修改`_config.yml`文件(有两个这样的文件,blog 目录下的是站点配置文件,next 主题目录下的是主题配置文件):

```shell
vim _config.yml
# 主题配置
theme: next
```

然后清除 hexo 缓存，重启服务(每次修改主题配置文件都要重启):

```shell
hexo clean
hexo s --debug
```

## 主题设定

### 选择 scheme

修改主题配置文件

```yml
# Schemes
#scheme: Muse
scheme: Mist
#scheme: Pisces
#scheme: Gemini
```

### 设置语言

在站点配置文件添加中文支持

```yml
language: zh-CN
```

### 设置菜单

修改主题配置文件:

```yml
menu:
  home: / || home
  about: /about/ || user
  tags: /tags/ || tags
  categories: /categories/ || th
  archives: /archives/ || archive
  # schedule: /schedule/ || calendar
  sitemap: /sitemap.xml || sitemap
  #commonweal: /404/ || heartbeat
```

上面的菜单项默认只提供了 home 和 archives，其他的需要自己创建。而且冒号前面的名称并不是直接显示在界面的，而是用于图标链接以及显示文本匹配，所以不能修改

下面修改菜单项的显示文本:

找到 `themes/next/languages/zh-CN.yml` , 然后就可以修改菜单对应的中文

```yml
menu:
  home: 首页
  archives: 归档
  categories: 分类
  tags: 标签
  about: 关于
  search: 搜索
  schedule: 日程表
  sitemap: 站点地图
  commonweal: 公益 404
```

#### 添加分类和标签

新建 categories 文件:

```shell
hexo new page categories
```

该文件下会默认生成 index.md 文件:

```markdown
---
title: categories
date: 2020-01-17 14:56:16
---
```

添加一行:

```markdown
type: categories
comments: false #这一行是为了禁用评论功能
```

后面创建博客时添加 categories 属性即可为文章分类,比如:

```markdown
---
title: categoriestest
date: 2020-01-17 14:41:52
categories: Hexo
---
```

同样的，添加标签页的步骤一样，只不过在对应的 index.md 中添加

```markdown
type: tags
comments: false #同上
```

然后为新创建的博客添加 tags 属性即可为文章添加标签:

```markdown
---
title: tagtest1
date: 2020-01-17 14:41:52
tags: 博客
---
```

### 设置侧栏

修改主题配置文件:

```yml
sidebar:
  # 侧栏位置
  position: left
  #position: right

  # 侧栏显示时机
  #  - post    expand on posts automatically. Default.
  #  - always  expand for all pages automatically.
  #  - hide    expand only when click on the sidebar toggle icon.
  #  - remove  totally remove sidebar including sidebar toggle.
  display: post
```

### 设置图像

主题配置文件

```yml
avatar:
  # Replace the default image and set the url here.
  url: images/avatar.jpg
  # If true, the avatar will be dispalyed in circle.
  rounded: false
  # If true, the avatar will be rotated with the cursor.
  rotated: false
```

其中 url 可以是

|       地址       | 值                                                                                                                                                                                   |
| :--------------: | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| 完整的互联网 URI | http://example.com/avatar.png                                                                                                                                                        |
|   站点内的地址   | 将头像放置主题目录下的 `source/uploads/` （新建 uploads 目录若不存在） 配置为：`avatar: /uploads/avatar.png`或者 放置在 `source/images/` 目录下 配置为：`avatar: /images/avatar.png` |

### 设置作者和站点描述

站点配置文件:

```yml
# Site
title: 半纸药笺
subtitle: ""
description: "变优秀!"
keywords:
author: Tiansir
language: zh-Hans
timezone: ""
```

## 主题美化

### 修改字体

修改主题配置文件:

```yml
font:
  enable: true
  host:
  global:
    external: true
    family: Monda
    size:
```

### 设置三角丝带背景

在 blog 目录下执行命令安装相关依赖:

```shell
git clone https://github.com/theme-next/theme-next-canvas-ribbon themes/next/source/lib/canvas_ribbon
```

修改主题配置文件:

```yml
canvas_ribbon:
  enable: true
  size: 90 # The width of the ribbon
  alpha: 0.6 # The transparency of the ribbon
  zIndex: -1 # The display level of the ribbon
```

### 设置加载进度条

在 blog 目录下执行以下命令安装依赖：

```shell
git clone https://github.com/theme-next/theme-next-pace themes/next/source/lib/pace
```

修改主题配置文件:

```yml
pace:
  enable: true
  # Themes list:
  # big-counter | bounce | barber-shop | center-atom | center-circle | center-radar | center-simple
  # corner-indicator | fill-left | flat-top | flash | loading-bar | mac-osx | material | minimal
  theme: center-atom
```

### 设置 GitHub 挂件

修改主题配置文件:

```yml
github_banner:
  enable: true
  permalink: https://github.com/Tiansir-wg
  title: Follow me on GitHub
```

### 设置建站时间

修改主题配置文件:

```yml
footer:
  # Specify the date when the site was setup. If not defined, current year will be used.
  since: 2017
```

### 添加访客统计

修改主题配置文件:

```yml
busuanzi_count:
  enable: true
  total_visitors: true
  total_visitors_icon: user
  total_views: true
  total_views_icon: eye
  post_views: true
  post_views_icon: eye
```

> 由于不蒜子是基于域名来进行统计计算的，所以通过 localhost:4000 端口访问的时候统计数据 UV 和 PV 都会异常的大，属于正常现象。

### 显示近期文章

修改主题配置文件:

```yml
recent_posts: true
recent_posts_layout: block
recent_posts_title: 近期文章
```

修改侧边栏布局(themes/next/layout/\_partials/sidebar/site-overview.swig):

```js
{%- if theme.links %}
  ...
{%- endif %}
//这里添加
{% if theme.recent_posts %}
    <div class="links-of-blogroll motion-element {{ "links-of-blogroll-" + theme.recent_posts_layout  }}">
      <div class="links-of-blogroll-title">
        <!-- modify icon to fire by szw -->
        <i class="fa fa-history fa-{{ theme.recent_posts_icon | lower }}" aria-hidden="true"></i>
        {{ theme.recent_posts_title }}
      </div>
      <ul class="links-of-blogroll-list">
        {% set posts = site.posts.sort('-date').toArray() %}
        {% for post in posts.slice('0', '5') %}
          <li>
            <a href="{{ url_for(post.path) }}" title="{{ post.title }}" target="_blank">{{ post.title }}</a>
          </li>
        {% endfor %}
      </ul>
    </div>
{% endif %}
```

### 自定义样式

在`themes/next/source/css`下创建`_custom`文件夹，然后创建``\_custom.styl`文件，样式写在里面即可，然后修改`themes/next/source/css/main.styl`文件，将该文件导入

```css
// 自定义样式
@import "_custom/_custom";
```

### 显示阅读进度

主题配置文件:

```yml
back2top:
  enable: true
  # Back to top in sidebar.
  sidebar: false
  # Scroll percent label in b2t button.
  scrollpercent: true
```

在自定义样式文件中添加返回顶部按钮的样式:

```css
// 回到顶部样式
.back-to-top .fa-arrow-up:before {
  color: rgb(
    random-color(0, 255) - 50%,
    random-color(0, 255) - 50%,
    random-color(0, 255) - 50%
  );
}
```

### 侧边栏目录展开

修改主题配置文件:

```yml
toc:
  # If true, all level of TOC in a post will be displayed, rather than the activated part of it.
  expand_all: true
```

## 内容优化

### 模板设置

在`scaffolds`文件夹下配置 draft.md 模板:

```markdown
---
title: { { title } }
categories:
tags:
date: { { date } }
---
```

### 文章发布修改时间

主题配置文件:

```yml
post_meta:
  item_text: true # 显示文字说明
  created_at: true # 显示文章创建时间
  updated_at:
    enable: false # 文章修改时间
    another_day: false # 只有当修改时间和创建时间不是同一天的时候才显示
  categories: true # 分类信息
```

### 文章字数统计

在 blog 目录下安装相关依赖:

```shell
cnpm install hexo-symbols-count-time --save
```

修改主题配置文件:

```yml
symbols_count_time:
  separated_meta: true # 是否换行显示 统计信息
  item_text_post: true # 文章统计信息中是否显示“本文字数/阅读时长”等描述文字
  item_text_total: false # 站点统计信息中是否显示“本文字数/阅读时长”等描述文字
```

然后重启即可

### 版权声明

主题配置文件:

```yml
creative_commons:
  license: by-nc-sa
  sidebar: true
  post: true
  language: zh-CN
```

### 修改链接样式

在`_custom.styl`文件下添加:

```css

$link-color = #2780e3;
$link-hover-color = #1094e8;
$sidebar-link-hover-color = #0593d3;

// 普通链接样式
a, span.exturl {
  &:hover {
    color: $link-hover-color;
    border-bottom-color: $link-hover-color;
  }
  // For spanned external links.
  cursor: pointer;
}

// 侧边栏链接样式
.sidebar a, .sidebar span.exturl{
  &:hover {
    color: $sidebar-link-hover-color;
    border-bottom-color: $sidebar-link-hover-color;
  }
}

// 侧边栏目录链接样式
.post-toc ol a {
  &:hover {
    color: $sidebar-link-hover-color;
    border-bottom-color: $sidebar-link-hover-color;
  }
}

//文章内链接文本样式
.post-body p a{
  color: $link-color;
  text-decoration: none;
  border-bottom: none;
  &:hover {
    color: $link-hover-color;
    text-decoration: underline;
    border-bottom-color: $link-hover-color;
  }
}

// 文章内上下一页链接样式
.post-nav-prev a , .post-nav-next a{
  &:hover {
    color: $link-hover-color;
  }
}
```

### 图片尺寸处理

hexo 不支持 Markdown 原生图片语法，无法设置图片大小，可以通过自定义图片尺寸处理脚本来解决

在`themes/next/source/js`文件夹下新建`custom.js`文件

```js
function set_image_size(image, width, height) {
  image.setAttribute("width", width + "px");
  image.setAttribute("height", height + "px");
}

function hexo_resize_image() {
  var imgs = document.getElementsByTagName("img");
  for (var i = imgs.length - 1; i >= 0; i--) {
    var img = imgs[i];
    var src = img.getAttribute("src").toString();
    var fields = src.match(/\?(\d*x\d*)/);
    if (fields && fields.length > 1) {
      var values = fields[1].split("x");
      if (values.length == 2) {
        var width = values[0];
        var height = values[1];

        if (!(width.length && height.length)) {
          var n_width = img.naturalWidth;
          var n_height = img.naturalHeight;
          if (width.length > 0) {
            height = (n_height * width) / n_width;
          }
          if (height.length > 0) {
            width = (n_width * height) / n_height;
          }
        }
        set_image_size(img, width, height);
      }
      continue;
    }

    fields = src.match(/\?(\d*)/);
    if (fields && fields.length > 1) {
      var scale = parseFloat(fields[1].toString());
      var width = (scale / 100.0) * img.naturalWidth;
      var height = (scale / 100.0) * img.naturalHeight;
      set_image_size(img, width, height);
    }
  }
}
window.onload = hexo_resize_image;
```

然后在`themes/next/layout/_scripts/index.swig`中添加:

```js
{
  {
    -next_js("custom.js");
  }
}
```

### 代码复制功能

在主题配置文件中开启即可:

```yml
copy_button:
  enable: true
  # Show text copy result.
  show_result: false
  # Available values: default | flat | mac
  style:
```

### 草稿和发布

> 使用`hexo new`来建立文章会将新文章建立在 **source/\_posts** 目录下，当使用 hexo generate 编译文件时，会将其 HTML 结果编译在 public 目录下，之后`hexo server`将会把 public 目录下所有文章发布。

hexo 提供草稿功能，其将文件建立在**source/\_drafts** 下

```
hexo new draft <title>	# 新建草稿文章
hexo s --draft	        # 预览草稿文章
```

下面命令可将草稿发布:

```shell
hexo P <filename>   #不包含 md 后缀的文章名称
```

## 功能强化

### 站内搜索

在 blog 目录下安装相关依赖:

```shell
cnpm install hexo-generator-searchdb --save
```

编辑站点配置文件，添加以下内容:

```yml
search:
  path: search.xml
  field: post
  format: html
  limit: 10000
```

修改主题配置文件:

```yml
local_search:
  enable: true
  # If auto, trigger search by changing input.
  # If manual, trigger search by pressing enter key or search button.
  trigger: auto
  # Show top n results per article, show all results by setting to -1
  top_n_per_article: 1
  # Unescape html strings to the readable one.
  unescape: true
  # Preload the search data when the page loads.
  preload: false
```

### 博客内容截断显示

在文章中使用 <!-- more --> 手动进行截断,也可以在文章首部添加 description 指定文章描述信息

### 评论功能

首先，在 LeanCloud 上注册账号并创建应用，设置 LeanCloud 的信息。

在 **存储** -> **数据** 中 新建一个名为`Counter` 的 Class，`ACL`权限设置为 **无限制**：
在 **设置** -> **安全中心** 中添加博客域名到 Web 安全域名中，以保护 LeanCloud 应用的数据安全。

修改主题配置文件:

```yml
valine:
  enable: true # 开启 Valine 评论
  # 设置应用 id 和 key
  appid: # your leancloud application appid
  appkey: # your leancloud application appkey
  # 关闭提醒与验证
  notify: false
  verify: false
  placeholder: # 文本框默认文字
  avatar: mm # gravatar style
  guest_info: nick,mail # 需要填写的信息字段
  pageSize: 10 # 每页评论数
  language: zh-cn # language, available values: en, zh-cn
  visitor: true # 开启文章阅读次数统计
  comment_count: false # 首页是否开启评论数
```

有时候我们并不想在文章标题下显示评论数量，如要隐藏，可在自定义样式文件中添加如下代码：

```css
//屏蔽标题下的评论数量
.post-comments-count {
  display: none;
}
```

评论区会显示评论人的浏览器和操作系统版本号等信息，如果只想要一个干净的评论界面，而没有多余其他的信息，可在自定义样式文件中添加如下代码：

```css
//屏蔽评论组件的多余信息
#comments .info,
#comments .vsys {
  display: none;
}
```

最后，集成评论服务后，所有的页面也会带有评论，包括标签、关于等页面。这里需要在添加字段`comments`并将值设置为 false 即可

```markdown
---
title: 标签
type: "tags"
comments: false
---
```

有关配置可以参考这篇[文章](https://xubuhui.coding.me/2019/07/22/Hexo-Valine-Leancloud%E5%BF%AB%E9%80%9F%E6%90%AD%E5%BB%BA%E8%AF%84%E8%AE%BA%E7%B3%BB%E7%BB%9F/)

### 图片沙箱

blog 目录下安装相关依赖:

```shell
git clone https://github.com/theme-next/theme-next-fancybox3 themes/next/source/lib/fancybox
```

修改主题配置文件:

```yml
fancybox: true
```

### 开启打赏功能

只需要在主题配置文件中填入支付宝和微信的收款二维码图片地址即可:

```yml
reward_settings:
  # If true, reward will be displayed in every article by default.
  enable: true
  animation: false
  comment: 坚持原创技术分享，您的支持将鼓励我继续创作！

reward:
  wechatpay: images/weichatpay.png
  alipay: images/zhifubao.jpeg
  #bitcoin: /images/bitcoin.png
```

### 友情链接

主题配置文件:

```yml
links:
  CSDN: https://blog.csdn.net/TMRsir
```

### 背景动画

主题配置文件:

```yml
# 下面两种只能开启一种
canvas_nest:
  enable: true
  onmobile: true # Display on mobile or not

three:
  enable: false
  three_waves: false
  canvas_lines: false
  canvas_sphere: false
```

### 数学公式支持

然后修改主题配置文件:

```yml
math:
  # Default (true) will load mathjax / katex script on demand.
  # That is it only render those page which has `mathjax: true` in Front-matter.
  # If you set it to false, it will load mathjax / katex srcipt EVERY PAGE.
  per_page: true

  # hexo-renderer-pandoc (or hexo-renderer-kramed) required for full MathJax support.
  mathjax:
    enable: true
    # See: https://mhchem.github.io/MathJax-mhchem/
    mhchem: true
```

后面如果需要在文章中使用数学公式的话要在 文章开头添加:

```markdown
---
mathjax: true
---
```

## 性能优化

### 设置 CDN 加速

修改主题配置文件的`vendors`配置，cdn 相关配置在文件中注释掉了，打开即可

```yml
vendors:
  # Internal path prefix.
  _internal: lib

  # Internal version: 3.1.0
  anime: //cdn.jsdelivr.net/npm/animejs@3.1.0/lib/anime.min.js
  #anime:
 ...
```

### 设置图床

参考我的另一篇文章[picgo+gitee 搭建图床](https://tiansir-wg.gitee.io/blog/2020/04/27/picgo-gitee搭建图床/)

## 部署到码云

> 部署之前需要有 git 并且配置好 ssh

安装 **hexo-deployer-git** :

```shell
cnpm install hexo-deployer-git --save
```

之所以部署到 码云而不是 github，是因为 github 国内访问实在太慢啦，本质上部署到两个平台的步骤是差不多的。

在部署到 gitee 之前需要先创建仓库，为了便于代码迁移可以设置两个分支，master 分支放置编译后的静态界面，另一个分支(develop 分支，设置为默认分支)放置源文件的配置

修改站点配置文件:

```yml
deploy:
  type: git
  repo: https://github.com/<username>/<project>
  # example, https://github.com/hexojs/hexojs.github.io
  branch: master
```

将仓库克隆到本地文件夹，然后复制 hexo 博客文件夹的文件到该文件夹，为了防止代码重复提交，可以将 public 下的文件添加到.gitignore 中，然后就 push 到仓库即可

```shell
git add .
git commit -m ""
git push

```

后面提交时可能会出现问题，这是因为 hexo 部署时会生成一个.deploy_git 文件，它也是个仓库，实际上提交源代码时是不需要它的，将其添加到.gitignore 文件即可

这一部分只是提交源文件，还需要在本地执行

```shell
hexo g
hexo d
```

这样编译后的静态文件就会提交到 master 分支，也就是站点配置文件中配置的分支

如果用 github 的话后面就完了，而用 gitee 的话还需要手动进行部署,然后才能访问


