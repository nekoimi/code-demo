# Swoole 的 Timer

### Timer

主要有三个方法

- swoole_timer_tick 间隔的时钟控制器

- swoole_timer_after 指定的时间后执行

- swoole_timer_clear 删除定时器

#### 本地环境

因为使用的windows, 所以跑在docker里面

- php 7.3.6

- swoole 4.3.3

### timer_tick & swoole_timer_after & swoole_timer_clear

```php
$timer_id = swoole_timer_tick(1000, function ($timer_id) {
    echo $timer_id . ' : ' . date('H:i:s') . PHP_EOL;
});

echo 'aaaaaaaaaaa' . PHP_EOL;

swoole_timer_after(3000, function () use ($timer_id){
    echo $timer_id . ' AAAAA : ' . date('H:i:s') . PHP_EOL;

    swoole_timer_clear($timer_id);
});
```

- 执行结果

![TIM截图20190619140813.png](https://i.loli.net/2019/06/19/5d09d155b08f977689.png)

也可以直接使用\Swoole\Timer类来创建不同timer应用

### 定时任务调度

![任务调度](https://raw.githubusercontent.com/osgochina/swoole-crontab/master/doc/xufei100.png)
