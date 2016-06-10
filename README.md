# Blog
* 使用Go语言，Beego框架开发，简单的实现个人Blog
* The Simple Blog base on Beego and using Golang

## 如何运行
### 1. 获取依赖
```
go get -u github.com/astaxie/beego
go get -u github.com/Unknwon/goconfig
go get -u github.com/garyburd/redigo/redis
```

### 2. 获取项目
```
go get -u github.com/RayFantasyStudio/blog
```

### 3. 配置数据库
本项目使用MySQL/MariaDb+Redis，确认安装好数据库Server后，在`<GOPATH>/github.com/RayFantasyStudio/blog/conf`中，新建名为`mysql.conf`的文件，写入以下内容
```
# 访问数据库的用户名
mysqluser = user

# 访问数据库的密码
mysqlpwd = password

# 数据库Server的主机名/IP地址/域名
mysqlhost = localhost

# 访问数据库Server的端口
mysqlport = 3306

# 数据库名
mysqldb   = rfs_blog
```
新建名为`redis.conf`的文件，写入以下内容
```
# 连接协议
redisnetwork = tcp

# 数据库Server地址
redisaddress = 127.0.0.1:6379
```
请将=后的内容改成你自己的

### 4. 编译运行
#### 直接编译运行
```
cd <GOPATH>/github.com/RayFantasyStudio/blog
go build main.go
./main
```

#### 使用bee工具
1. 安装bee工具
```
go get github.com/beego/bee
```

2. 通过bee工具编译运行
```
cd <GOPATH>/github.com/RayFantasyStudio/blog
bee run
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
