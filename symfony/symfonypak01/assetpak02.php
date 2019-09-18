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

require_once __DIR__ . '/../vendor/autoload.php';
use Symfony\Component\Asset\UrlPackage;
use Symfony\Component\Asset\VersionStrategy\StaticVersionStrategy;

$asset = new UrlPackage(
    '//www.baidu.com',
    new StaticVersionStrategy('v1', '%2$s/%1$s')
);

echo $asset->getUrl('style.css');
# //www.baidu.com/v1/style.css
