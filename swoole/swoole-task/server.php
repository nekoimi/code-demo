<?php

class Server {

    /**
     *
     * @var \Swoole\Server
     */
    private $server;

    /**
     * Events
     */
    private $events = array(
        'start',
        'connect',
        'receive',
        'close',
        'task',
        'finish'
    );

    /**
     * Server constructor.
     */
    public function __construct()
    {
        $this->server = new \Swoole\Server('0.0.0.0', 9801);
        $this->server->set(array(
            'worker_num'        => 2,   // 开启2个worker进程
            'max_request'       => 4,   // 这里是每个worker进程的 max_request 为4
            'task_worker_num'   => 4,   // 开启4个task进程
            'dispatch_mode'     => 2,   // 数据分发策略 -> 固定模式  @see https://wiki.swoole.com/wiki/page/277.html
        ));
        foreach ($this->events as $event) {
            if (method_exists($this, 'on' . ucfirst($event))) {
                $this->server->on($event, [$this, 'on' . ucfirst($event)]);
            }
        }
        $this->server->start();
    }

    // start
    public function onStart (\Swoole\Server $server) {
        echo __METHOD__ . PHP_EOL;
    }

    // connect
    public function onConnect (\Swoole\Server $server, $fd) {
        echo __METHOD__ . PHP_EOL;
    }

    // receive
    public function onReceive (\Swoole\Server $server, $fd, $from_id, $data) {
        echo __METHOD__ . PHP_EOL;
    }

    // task
    public function onTask (\Swoole\Server $server, $fd, $from_id, $data) {
        echo __METHOD__ . PHP_EOL;
    }

    // finish
    public function onFinish (\Swoole\Server $server, $task_id, $data) {
        echo __METHOD__ . PHP_EOL;
    }

    // close
    public function onClose (\Swoole\Server $server, $fd) {
        echo __METHOD__ . PHP_EOL;
    }

}