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
 * #     19-6-2  下午1:58
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/

$f = fopen(__DIR__ . '/test_file.php', 'r');
\Swoole\Coroutine::create(function () use ($f) {
    $res = co::fread($f);
    var_dump($res);
});

go(function () {
    var_dump(
      co::readFile(__DIR__. '/test_file.php')
    );
});

echo 'master' . PHP_EOL;
