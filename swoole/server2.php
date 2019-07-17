<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-17 下午3:33
 * #                            ------
 **/
$pid = pcntl_fork();

if ($pid == -1) {
    die('could not fork.');
} elseif ($pid) {
    var_dump('父进程', $pid);
    pcntl_wait($status);
}else {
    // 子进程执行代码的区域
    var_dump('子进程', $pid);
    sleep(10);
}