<?php
/**
 * # ----
 * #     Yprisoner <yyprisoner@gmail.com>
 * #                   19-7-10 下午3:39
 * #                            ------
 **/
require_once '../vendor/autoload.php';
use Symfony\Component\Cache\Adapter\FilesystemAdapter;
$cache = new FilesystemAdapter();

$numProducts = $cache->getItem('status.num_products');
$numProducts->set(233333);
$cache->save($numProducts);


