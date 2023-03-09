# 使用指引

## 新增文章

切换到`markdown`分支，执行`hugo new post xxx`生成新文章。编辑完成后发不到markdown分支。

```shell
git add .
git commit -m "新增文章"
git push origin markdown
```

## 发布

源码文件发布到`markdown`分支，编译后的文件发布到`master`分支。

```shell
git checkout markdown
hugo -D
pushd public
git add .
git commit -m "新增文章"
git push origin master
popd

```



