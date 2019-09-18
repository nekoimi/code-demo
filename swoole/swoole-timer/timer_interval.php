<?php

/**
 * 间隔的时间 -> 跟 js 里面的 setInterval 函数一样
 *
 * 可以用来代替 linux 的 Crontab  [ 分 时 日 月 周 ]
 *
 * Swoole 却可以精确到毫秒!
 */

//$timer = \Swoole\Timer::tick(3000, function () use ($timer) {
//    echo time() . PHP_EOL;
//});
$exec_num = 10;
$timer_id = swoole_timer_tick(3000, function () use ($timer_id, &$exec_num) {
    if ($exec_num === 0) {
        swoole_timer_clear($timer_id); // stop
    }
    echo  'hello world.  ' . strval(time()) . PHP_EOL;
    $exec_num --;
});
