<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-21 下午2:13
 * #                            ------
 **/
namespace Service;

use Swoole\Server;

abstract class BaseServer implements ServerInterface {

    public function onStart(Server $server) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onShutdown(Server $server) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onWorkerStart(Server $server, int $worker_id) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onWorkerStop(Server $server, int $worker_id) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onWorkerExit(Server $server, int $worker_id) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onConnect(Server $server, int $fd, int $reactorId) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onReceive(Server $server, int $fd, int $reactorId, string $data) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }

    public function onClose(Server $server, int $fd, int $reactorId) {
        echo __CLASS__ . __METHOD__ . PHP_EOL;
    }
}