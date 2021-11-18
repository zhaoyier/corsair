## 准备工作

### Node 版本 > 10

-   nvm alias default 13
-   react 在最新的 16.8 中启用了 hooks 语法，力推函数式组件，尽管官方表示 class 式的组件在后续版本中并不会废弃，但是 hooks 是未来前端框架中组件的发展方向（最新的 Vue 也借鉴了 react Hook 的很多思路），我们需要大胆尝试新鲜事物

### 克隆相关仓库

    $ git clone ssh://git@git.ezbuy.me:10022/ezbuy/ezship.git  // ezship 开发环境仓库
    $ git clone ssh://git@git.ezbuy.me:10022/deploy/ezship.65emall.net.git  // ezship uat环境源码仓库
    $ git clone ssh://git@git.ezbuy.me:10022/deploy/ezship.65daigou.com.git // ezship 生产环境源码仓库

`ezship`, `ezship.65emall.net`, `ezship.ezbuy.com` 确保为同级目录。

    $ ls -d \*/
    $ ezship
    $ ezship.65emall.net
    $ ezship.65daigou.com

## 开发

```
$ make dev
// 或
$ yarn start
// kflow本地环境
$ yarn start@kflow
```

## 发布 UAT 环境

在 **开发分支** 执行 `make publish`, 将自动视为发布 UAT 环境。

    $ git checkout <branch>
    $ make publish
    $ ...

## 生产环境部署

在 **master** 执行`make publish`，将自动视为发布至生产环境。

    $ git checkout master
    $ make publish
    $ ...

## 增加新页面

在 `src/views` 的目录下，新建一个目录，然后运行 `make genEntry`。目录名称就是页面的地址，比如新建的目录是 `helloWorld`，那访问的路径就是 `/helloWorld.html`。

## 目前接口配置

-   登录 admindoc/AdminLogin.thrift Login
