<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-17 下午3:39
 * #                            ------
 **/
$s = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
socket_set_option($s, SOL_SOCKET, SO_REUSEADDR, 1);
socket_bind($s, '0.0.0.0', 9090);
socket_listen($s, 1);
while ($client = socket_accept($s)) {
    $pid = pcntl_fork();
    if ($pid == -1) {
        die('fork error');
    } elseif ($pid) {
        var_dump('todo 父进程');
    } else {
        sleep(10);
//        $buffer = socket_read($client, 1024);
//        var_dump($buffer);
        $string = "<h1>Hello World.</h1>";
        $header = "HTTP/1.1 200 OK\r\n";
        $header .= "Content-Type: text/html;charset=utf-8\r\n\r\n";
        socket_write($client, $header . $string);
        socket_close($client);
        exit();
    }
}
socket_close($s);