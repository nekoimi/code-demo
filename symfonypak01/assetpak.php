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

use Symfony\Component\Asset\Package;
use Symfony\Component\Asset\VersionStrategy\StaticVersionStrategy;

$asset = new Package(
    new StaticVersionStrategy('v1', '%2$s/%1$s')
);

echo $asset->getUrl('style.css');
// v1/style.css