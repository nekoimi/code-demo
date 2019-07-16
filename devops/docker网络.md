# Docker 网络 ==> Bridge


docker 默认的虚拟网卡 `docker0`

所有的网络都使用docker0网络， 除非你在使用docker run --net指定其他选项

![20190716131052.png](https://i.loli.net/2019/07/16/5d2d5c7d1e53a87181.png)

所有的docker容器都使用这个docker0的网卡

查看docker的网络

```bash
docker network list
```

![深度截图_选择区域_20190716130612.png](https://i.loli.net/2019/07/16/5d2d5b66dbe8338090.png)


可以查看各个网络的详情

```bash
docker network inspect 91aa6b806110
```

### 两个容器之间的通信
 
先跑两个容器：

![20190716132106.png](https://i.loli.net/2019/07/16/5d2d5f2a4a9cb41394.png)

![20190716132155.png](https://i.loli.net/2019/07/16/5d2d5f3f4423791749.png)
 
在一个Docker容器内部可以访问其他的容器地址

链接都是基于 docker 
 
### 容器如何访问外网

在TCP/IP中，我们知道，一个主机访问互联网，发送数据和接受数据都需要经过网络模型中的几层模型

再经过链路层的时候需要将逻辑地址转换成本地主机的网卡物理地址

这里使用的就是ARP（网络地址转换）
