<?php

class WebSocketServer {

    /**
     *
     * @var \Swoole\WebSocket\Server
     */
    private $server;

    /**
     * WebSocketServer constructor.
     */
    public function __construct()
    {
        $this->server = new \Swoole\WebSocket\Server('0.0.0.0', 9802);
        $this->server->set(array(
            'worker_num' => 2, // 开启2个worker进程
            'max_request' => 4,
            'task_worker_num' => 4, // 开启4个task进程
            'dispatch_mode' => 4, // 数据包分发策略 - IP分配
            'daemonize' => false, // 守护进程
        ));
    }

    public function onTask (\Swoole\WebSocket\Server $server, $task_id, $from_id, $data) {
        foreach ($this->server->connections as $connection) {
            $this->server->push($connection, json_encode(array(
                'message' => 'demo100'
            )));
            // 这里的 $data 数据最大长度不超过2M
        }
        $this->server->finish($data);
    }

}
