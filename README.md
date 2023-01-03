<h align="center">如何把html、js和css等等静态文件打包成各个平台可执行文件?</h>

## step 1 把这个项目fork到自己的仓库中
然后下载下来

## step 2 把要打包的静态文件放到assets/目录下 

``` bash
.
├── Makefile
├── README.md
├── assets # <-- 就是放在这个目录下
├── dist
├── go.mod
├── go.sum
└── main.go
```

## step 3 提交代码
``` bash 
$ git add -A 
$ git commit -m 'feat(static files): 更新静态文件'
$ git push origin master
$ git tag v0.0.3 # <-- 要打上新的版本号，这样才能触发github action的构建
$ git push --tags # <-- 把新版本推送上去

```

## step 4 查看github的构建流程

把以下链接的用户名和仓库名修改为你自己fork的，然后就能看见自己的打包构建流程了
https://github.com/wuchuheng/com.wuchuheng.go.serverForStaticFile/actions


## step 5 查看构建完成，已经打包好的个平台的文件
把以下链接的仓库名换成你fork过去的仓库名，然后打开，就能看到你自己在github action中打包好的各个平台的可执行文件了。如: 
https://github.com/wuchuheng/com.wuchuheng.go.serverForStaticFile/releases
你要找的window平台的文件就是以`exe`结尾的文件，下载它然后，在命令行下启动它就可以了。






