<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-21 下午1:37
 * #                            ------
 **/
namespace Service;

use Swoole\Server;

interface ServerInterface {

    public function onStart(Server $server);
    public function onShutdown(Server $server);
    public function onWorkerStart(Server $server, int $worker_id);
    public function onWorkerStop(Server $server, int $worker_id);
    public function onWorkerExit(Server $server, int $worker_id);
    public function onConnect(Server $server, int $fd, int $reactorId);
    public function onReceive(Server $server, int $fd, int $reactorId, string $data);
    public function onClose(Server $server, int $fd, int $reactorId);

}