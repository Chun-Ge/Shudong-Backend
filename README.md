# Shudong-Backend

中大树洞 Backend Git Repo

使用Go的[Iris](https://github.com/kataras/iris)框架开发， 官方用例在此： [Link](https://iris-go.com/v10/recipe)

简版代码规范：见[backend-code-standard in Chun-Ge/documents](https://github.com/Chun-Ge/documents/blob/master/technical-docs/backend-code-standard.md)

[Effective Go 中英文对照版](https://legacy.gitbook.com/book/bingohuang/effective-go-zh-en/details)

## Go 使用注意事项

- 安装`golang.org`下的包时可能出现网络错误(原因都懂)，现提供几种解决方法
  - 在命令行下科学上网
  - 参考[这篇文章](https://studygolang.com/articles/10263)
  - 如果在中大校园网内，可以使用[LGA1150](https://github.com/LGA1150)维护的[SYSUv6-DNS](https://github.com/bazingaterry/SYSUv6-DNS)
- 自身代码的包引用请务必采用[vendor模式](https://www.jianshu.com/p/e52e3e1ad1c0)，否则合作开发时会导致非常大的麻烦！
  - 同时也不要使用如`import .`或 `import ../xxx`的相对路径方式引用包
- 推荐使用vscode进行go的开发，插件齐全，简单易用。。

## 开发时的注意事项

- 使用`dev`分支作为开发主分支，`master`作为稳定版分支。`dev`分支下的内容确定运行正常且基本无需改动之后，可以合并到`master`分支下，作为当前的稳定版。
- 请大家开发过程中**不要**直接使用`master`或`dev`开发
- 分支命名请使用 feature/**** (你的特性名称)， 如 feature/create-post 表示实现了新的feature
- `master`或`dev`出现bug时，从`master`或`dev`切出分支，命名为 hotfix/**** (修复的bug) (别忘了必要的rebase)
- 每完成一个分支请用发PR的形式让大家一起讨论， OK之后再进行合并

## TODO

- ~~确定并建立文件结构并更新Readme: Iris不规定明确的文件结构~~
- ~~与前端讨论更明确的API~~
- 一阶段开发目标：~~登入/登出、注册、~~ **发帖、回复、**~~点赞(帖子、回复)~~
- 二阶段开发目标：
  - 功能：收藏(帖子)、举报(帖子、评论)、密码(修改、找回)、删除(帖子、评论；仅限作者自删)
    - Low priority: 分享(shareCount++, 其他交给前端)
  - middlewares.InternalErrorCatcher
  - iris.Configuration
  - Logger
  - go test

## License

No License
