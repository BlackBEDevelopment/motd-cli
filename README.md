# motd-cli
像PING一样获取我的世界服务器MOTD信息

基于[MCBE-Server-Motd](https://github.com/BlackBEDevelopment/MCBE-Server-Motd)项目开发的MOTD命令行版本

## 🛫 如何使用
先前往[releases](https://github.com/BlackBEDevelopment/motd-cli/releases)下载最新版本解压即可
### 命令参数
```
motd.exe be nyan.xyz:19132
```
第一个参数为服务器类型，可选`be`, `je`  
第二个参数为服务器地址，如果不指定端口则会尝试默认端口
### 推断模式
你也可以直接在第一个参数输入服务器地址，程序会自动尝试推断服务器类型(这会可能会导致获取信息的时间变长)  
例如: 
```
motd.exe nyan.xyz:19132
```

## ⚙️ 构建程序
自行构建前需要拥有 Go >= 1.18等必要依赖
### 克隆仓库
```
git clone https://github.com/BlackBEDevelopment/motd-cli.git
```
### 获取依赖
``` shell
go mod tidy
```
### 编译项目(Windows)
``` shell
go build -o motd.exe
```
### 编译项目(Linux)
``` shell
go build -o motd
```

## 🎬 设置为环境变量后使用
为了获得更好的体验，你可以尝试设置环境变量来通过如下方式在命令行使用程序
```
motd nyan.xyz
```
### Windows
1. 将下载到的文件解压到单独的文件夹内，记住文件夹的路径
2. 打开`高级系统设置`
3. 在高级选项卡选择环境变量
4. `用户环境变量`选项框内找到`Path`选项
5. 打开`Path`选项，选择`新建`，然后选择`浏览`后选择刚刚解压的程序的文件夹
6. 点击`确定`后重启电脑即可使用
### Linux
1. 将下载到的文件解压到单独的文件夹内，记住文件夹的路径
2. 输入`vim ~/.bashrc`打开环境变量配置
3. 在最后一行加上`export PATH=$PATH:{{path}}` {{path}}为刚刚存放程序的文件夹路径
4. 保存文件后输入`source ~/.bashrc`使配置文件生效即可