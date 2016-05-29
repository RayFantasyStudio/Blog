# Blog
* 使用Go语言，Beego框架开发，简单的实现个人Blog
* The Simple Blog base on Beego and using Golang

## 如何运行
### 1. 导入项目
```
go get github.com/astaxie/beego
go get github.com/RayFantasyStudio/blog
```

### 2. 配置数据库
目前仅支持MySQL/MariaDb，确认安装好SQL Server后，在`<GOPATH>/github.com/RayFantasyStudio/blog/conf`中，新建名为`mysql.conf`的文件，写入以下内容
```
# 访问数据库的用户名
mysqluser = "user"

# 访问数据库的密码
mysqlpass = "password"

# SQL Server的主机名/IP地址/域名
mysqlurls = "localhost"

# 访问SQL Server的端口
mysqlport = 3306

# 数据库名
mysqldb   = "rfs_blog"
```
请将=后的内容改成你自己的

### 3. 编译运行
```
cd <GOPATH>/github.com/RayFantasyStudio/blog
go build main.go
./main
```

## 许可证(License)
```
Copyright 2015-2016 RayFantasy Studio

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
```
