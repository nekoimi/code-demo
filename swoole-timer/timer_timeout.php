<?php

/**
 * 延迟执行任务
 */

$now = time();
$timer = swoole_timer_after(3000, function () use ($now) {
    echo time() - $now;
});
