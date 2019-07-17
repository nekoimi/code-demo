<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-17 下午3:31
 * #                            ------
 **/
$server = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
socket_bind($server, '0.0.0.0', 9090);
socket_listen($server, 1);
socket_set_option($server, SOL_SOCKET, SO_REUSEADDR, 1);  // 复用链接
while (true) {
    $client = socket_accept($server);
    $buffer = socket_read($client, 1024);
    var_dump($buffer);
    $string = "<h1>Hello World.</h1>";
    $header = "HTTP/1.1 200 OK\r\n";
    $header .= "Content-Type: text/html;charset=utf-8\r\n\r\n";
    socket_write($client, $header . $string);
    socket_close($client);
}
socket_close($server);