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
 * #     19-6-23  下午10:07
 * #                            ------
 * #    「 涙の雨が頬をたたくたびに美しく 」
 **/
$databaseConfig = array(
    'driver'    => 'mysql',
    'host'      => '127.0.0.1',
    'port'      => 3306,
    'database'  => 'demo',
    'username'  => 'root',
    'password'  => 'root',
    'charset'   => 'utf8mb4',
    'collation' => 'utf8mb4_unicode_ci',
    'prefix'    => '',
    'timezone'  => '+08:00',
    'strict'    => true,
);

require_once '../../vendor/autoload.php';

$connectionFactory = new \Illuminate\Database\Connectors\ConnectionFactory(
    new \Illuminate\Container\Container()
);
$connection = $connectionFactory->make($databaseConfig);


// 添加
$insertId = $connection->table('users')->insertGetId(array(
    'username' => '张三',
    'password'  =>  password_hash('demo100', PASSWORD_DEFAULT)
));
print_r($insertId);  // 返回最后一次添加的自增ID
// 修改
$result = $connection->table('users')->where('username', '张三')->update(array(
    'password'  =>  password_hash('demo200', PASSWORD_DEFAULT)
));
print_r($result);  // 返回受影响行数
// 删除
$result = $connection->table('users')->where('username', '张三')->delete();
print_r($result);  // 返回受影响行数
// 查询
$users = $connection->table('users')->get();
print_r($users);  // Collection
