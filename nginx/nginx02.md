# Nginx 入门总结 - Nginx基础

> Nginx是一款轻量级的Web 服务器/反向代理服务器及电子邮件（IMAP/POP3）代理服务器，在BSD-like 协议下发行。其特点是占有内存少，并发能力强，事实上nginx的并发能力确实在同类型的网页服务器中表现较好，中国大陆使用nginx网站用户有：百度、京东、新浪、网易、腾讯、淘宝等。

### 包管理器安装

- ArchLinux 

```bash
sudo pacman -S nginx
```

- Debian系

```bash
apt install -y nginx
```

- RedHat系 (Centos、Fedora)

```bash
yum install -y nginx
```

### 编译安装

- 需要自行安装相关的lib

<a href="https://nginx.org/en/download.html" target="_blank">下载Nginx</a>

```bash
wget 下载地址

tar -zxvf nginx.1.16.tar.gz

cd nginx.1.16

./configure

make && make install
```

### Nginx的命令行

- 基本格式:    nginx  -s <command> 

- 查看帮助:    nginx  -? (-h) 

- 使用指定的conf文件启动:    nginx  -c  <.conf file path> 

- 测试配置文件是否有语法错误:    nginx  -t

- 重新打开日志文件:    nginx  -s   reopen

- 获取nginx的编译信息:    nginx  -v


### 基本配置

头部配置

![624133001.png](https://i.loli.net/2019/06/24/5d105ff0ceb8962027.png)

基本配置

![624133012.png](https://i.loli.net/2019/06/24/5d105ff0dab9b25969.png)

#### 内置变量

![624133145.png](https://i.loli.net/2019/06/24/5d10606b32ecb63872.png)
![624133201.png](https://i.loli.net/2019/06/24/5d10606b24b9e11570.png)

### Nginx Sendfile 零 Copy

一次文件读写的过程，由于用户空间并没有权限去操作底层的硬件设备，所以会发生上下文切换，

![624133353.png](https://i.loli.net/2019/06/24/5d10611f18bca26743.png)

sendfile 零拷贝 其实是linux系统底层的传输过程

一个文件的读取需要经过4次的copy过程

1.从硬盘copy到内核空间

2.从内核中间copy到用户空间

3.从用户空间copy到内核缓冲区

4.将缓冲区的数据cpoy到输出端(socket)

使用 sendfile 可以减少copy次数
直接通过内核空间传输文件

