<?php
/**
 * # ------------Oooo---
 * # -----------(----)---
 * # ------------)--/----
 * # ------------(_/-
 * # ----oooO----
 * # ----(---)----
 * # -----\--(--
 * # ------\_)-
 * # ----
 * #     author : Yprisoner <yyprisoner@gmail.com>
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/

$timer_id = swoole_timer_tick(1000, function ($timer_id) {
    echo $timer_id . ' : ' . date('H:i:s') . PHP_EOL;
});

echo 'aaaaaaaaaaa' . PHP_EOL;

swoole_timer_after(3000, function () use ($timer_id){
    echo $timer_id . ' AAAAA : ' . date('H:i:s') . PHP_EOL;

    swoole_timer_clear($timer_id);
});
