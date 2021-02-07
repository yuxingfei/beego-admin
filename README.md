# beego-admin 通用后台系统

#### [beego-admin](https://github.com/yuxingfei/beego-admin) v2.0.1版本，基于beego 2.0.1 框架和AdminLte前端框架，开发的go语言通用后台系统，在beego框架的基础上，封装了后台系统的分页功能，excel数据导出功能等丰富常用的扩展，基于MVC模式，html界面随心定义，相较于某些后台复杂代码生成的前端html元素，使用原生的html原生作为前端显示，更加的灵活自由。[beego-admin](https://github.com/yuxingfei/beego-admin)通用的后台系统真正的做到了开箱即用，欢迎大家使用。技术交流群：1151174994

## beego-admin 安装

### 安装方式 (GO MOD方式安装,已移除 GOPATH方式安装说明，需要的请查看 tag v1.0.1)

#### 1、安装beego v2.0.1和bee v2.0.2
参考[Beego](https://beego.me/docs/install/)和[Bee](https://beego.me/docs/install/bee.md)安装手册

#### 2、clone 项目到本地 GOPATH src目录之外的路径下
```
GitHub:   git clone git@github.com:yuxingfei/beego-admin.git
```
或
```
码云:   git clone git@gitee.com:yuxingfei/beego-admin.git
```


#### 3、配置数据库
```
将目录中beego-admin.sql文件导入mysql数据库

更改根目录下的config.yaml文件内的数据库连接信息
```

#### 4、安装项目依赖
```
beego-admin目录下 go mod tidy 将自动下载依赖包
```

### 通过上面方式安装后,接下来

#### 运行系统
```
直接运行go run main.go，或者使用bee run在项目下运行，开始进行调试开发
```

#### 访问后台
访问`/admin/index/index`，默认超级管理员的账号密码都为`super_admin`。


## 补充
[beego-admin](https://github.com/yuxingfei/beego-admin) 项目在beego v2.0.1的框架基础上完善了很多丰富的常用后台功能，分页封装、excel数据一键导出等功能，目前没有开发手册，相信大家一看代码就可知道功能怎么使用，如果大家需要详细的使用手册，我可为大家写一份详细的功能使用介绍，此外，如果有需要php语言的laravel版本的后台管理系统，可以使用[laravel-admin](https://github.com/yuxingfei/laravel-admin)。

## 注意！！！
当前最新master版本beego v2.0.1框架的版本，如果需要beego1.x 版本的请下载 tag v1.0.1 版本，因 Beego 2.x 的XSRF只支持 HTTPS 协议，所有app.conf配置中默认关闭了XSRF安全过滤，如有需要请手动开启,因beego v2.X版本和 beego v1.x版本区别较大，请根据beego最新手册进行开发使用。

技术交流QQ群：1151174994

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-admin-qq-share.png)

#### [beego-admin](https://github.com/yuxingfei/beego-admin)通用后台系统效果图

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-login-1.jpg)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-index-2.png)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-user-index-3.png)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-user-4.png)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-admin-role-5.png)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-admin-menu-6.png)

![Image](https://raw.githubusercontent.com/yuxingfei/images/master/beego-setting-7.png)
