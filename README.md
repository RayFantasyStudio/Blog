# Blog
* 使用Go语言，Beego框架开发，简单的实现个人Blog
* The Simple Blog base on Beego and using Golang
* 项目预览 Preview [RayFantasyStudio Blog](http://blog.rayfantasy.com/)

## 如何运行(How to run)
### 1. 获取依赖(Dependencies)
```
go get -u github.com/astaxie/beego
go get -u github.com/Unknwon/goconfig
go get -u github.com/garyburd/redigo/redis
```

### 2. 获取项目(Get blog)
```
go get -u github.com/RayFantasyStudio/blog
```

### 3. 配置数据库(Configuration database)
本项目使用MySQL/MariaDb+Redis，确认安装好数据库Server后，在`<GOPATH>/github.com/RayFantasyStudio/blog/conf`中，新建名为`mysql.conf`的文件，写入以下内容

This project use MySQL/MariaDb+Redis.create the file in `<GOPATH>/github.com/RayFantasyStudio/blog/conf` `mysql.conf`and follow this config
```
# 访问数据库的用户名(DatabaseUser)
mysqluser = user

# 访问数据库的密码(Pwd)
mysqlpwd = password

# 数据库Server的主机名/IP地址/域名(Host)
mysqlhost = localhost

# 访问数据库Server的端口(Port)
mysqlport = 3306

# 数据库名(DatabaseName)
mysqldb   = rfs_blog
```
新建名为`redis.conf`的文件，写入以下内容

create file  `redis.conf` at the same folder and follow this config
```
# 连接协议(Connection Agreement)
redisnetwork = tcp

# 数据库Server地址(Redis host)
redisaddress = 127.0.0.1:6379
```
请将=后的内容改成你自己的


### 4. 编译运行(Run)
#### 直接编译运行(Directly compile and run)
```
cd <GOPATH>/github.com/RayFantasyStudio/blog
go build main.go
./main
```

#### 使用bee工具(Use bee tool)
1. 安装bee工具(install bee tool)
```
go get github.com/beego/bee
```

2. 通过bee工具编译运行(use bee tool)
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
