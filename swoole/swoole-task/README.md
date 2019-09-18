- 异步Task 实现调用异步任务的执行

- Master 进程

- Worker 进程

- TaskWorker 进程

### 数据分发策略

- 1，轮循模式，收到会轮循分配给每一个Worker进程

- 2，固定模式，根据连接的文件描述符分配Worker。这样可以保证同一个连接发来的数据只会被同一个Worker处理

- 3，抢占模式，主进程会根据Worker的忙闲状态选择投递，只会投递给处于闲置状态的Worker

- 4，IP分配，根据客户端IP进行取模hash，分配给一个固定的Worker进程。可以保证同一个来源IP的连接数据总会被分配到同一个Worker进程。算法为 ip2long(ClientIP) % worker_num

- 5，UID分配，需要用户代码中调用 Server->bind() 将一个连接绑定1个uid。然后底层根据UID的值分配到不同的Worker进程。算法为 UID % worker_num，如果需要使用字符串作为UID，可以使用crc32(UID_STRING)

- 7，stream模式，空闲的Worker会accept连接，并接受Reactor的新请求
