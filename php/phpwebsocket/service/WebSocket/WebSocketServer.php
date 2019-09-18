<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-21 下午2:16
 * #                            ------
 **/
namespace Service\WebSocket;
use Service\BaseServer;
use Swoole\Http\Request;
use Swoole\WebSocket\Frame;
use Swoole\WebSocket\Server;

class WebSocketServer extends BaseServer implements WebSocketInterface {

    /**
     * @var Server
     */
    private $ws;

    public function __construct() {

    }

    public function start() {
        $this->ws->start();
    }

    public function onOpen(Server $server, Request $request) {
        // TODO: Implement onOpen() method.
    }

    public function onMessage(Server $server, Frame $frame) {
        // TODO: Implement onMessage() method.
    }

}