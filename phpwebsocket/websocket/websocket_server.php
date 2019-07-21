<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-21 上午9:41
 * #                            ------
 **/ 
$server = new \Swoole\WebSocket\Server('0.0.0.0', 8002);
$table = new \Swoole\Table(1024);
$table->column('uid', \Swoole\Table::TYPE_STRING, 32);
$table->column('fd', \Swoole\Table::TYPE_INT);
$table->create();

$server->table = $table;

/**
 *      *  事件列表
 *
 *  * onStart
 *  * onShutdown
 *  * onWorkerStart
 *  * onWorkerStop
 *  * onTimer
 *  * onConnect
 *  * onReceive
 *  * onClose
 *  * onTask
 *  * onFinish
 *  * onPipeMessage
 *  * onWorkerError
 *  * onManagerStart
 *  * onManagerStop
 *  WebSocket
 *  * onOpen
 *  * onHandshake
 *  * onMessage
 */

$server->on('open', function (\Swoole\WebSocket\Server $server, \Swoole\Http\Request $request) use ($table) {
    echo 'WebSocket open.' . PHP_EOL;
    var_dump($table);
    if (isset($request->get['uid'])) {
        $uid = $request->get['uid'];
        $server->table->set($uid, array (
            'uid'   =>  $uid,
            'fd'    =>  $request->fd
        ));
//        var_dump($server->table);
    }
});

$server->on('message', function (\Swoole\WebSocket\Server $server, \Swoole\WebSocket\Frame $frame) {
    echo 'WebSocket message.' . PHP_EOL;
    $msg = json_decode($frame->data, true);
    print_r($msg);
    $user = $server->table->get($msg['uid']);
    print_r($user);
    if ($user && $server->isEstablished($user['fd'])) {
        // 向客户端推送消息 -> 长度不超过2M
        $server->push($user['fd'], $msg['message']);
    }
});

$server->on('close', function ($ser, $fd) {
    echo "server close.   ===> {$fd}" . PHP_EOL;
//    $table->del();
});

$server->start();