### 进程

- 可以使用`ps`命令查看正在运行的进程

临时使用windows, 在docker跑，发现没有基本的命令

通过一下命令安装

```bash
apt install -y ps-watcher
```

查看CPU信息

```bahs
cat /proc/cpuinfo | grep "model name"
```

查看CPU个数

```bash
cat /proc/cpuinfo | grep "model name" | wc -l
```

### 平均负载

```bash
uptime
```

> 平均负载是指单位时间内，系统处于可运行状态和不可中断状态的平均进程 数，也就是平均活跃进程数

跟cpu的使用率没有直接的关系

可运行状态:

> 是指正在使用CPU或者正在等待CPU的进程，通过`ps`命令看到的处于`R`(Running / Runnable)状态的进程

不可中断进程:

> 是指正处于内核态关键流程中的进程，这些流程不可以被打断，通过`ps`命令看到的`D`(Uninterruptible Sleep, 也称为 Disk Sleep)状态的进程
>
> 不可中断状态实际上是系统对进程和硬件设备的一种保护机制

### 理想情况

> 平均负载最理想的情况是等于 CPU 个数

- CPU 密集型进程，使用大量 CPU 会导致平均负载升高
- I/O 密集型进程，等待 I/O 也会导致平均负载升高，但 CPU 使用率不一定很高
