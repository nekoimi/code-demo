<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-21 下午2:10
 * #                            ------
 **/
namespace Service\WebSocket;
use Service\ServerInterface;
use Swoole\Http\Request;
use Swoole\WebSocket\Frame;
use Swoole\WebSocket\Server;

interface WebSocketInterface extends ServerInterface {

    public function start();
    public function onOpen(Server $server, Request $request);
    public function onMessage(Server $server, Frame $frame);

}