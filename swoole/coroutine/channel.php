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
 * #     19-6-2  下午1:35
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/

/**
 * channel  跟go语言很相似
 *
 * - 支持多生产者协程和多消费者协程
 */
$chan = new Chan();

go(function () use ($chan) {
    while (true) {
        $chan->push(time());
        co::sleep(1);
    }
});

go(function () use ($chan) {
    while (true) {
        $data = $chan->pop();
        echo $data . PHP_EOL;
    }
});
