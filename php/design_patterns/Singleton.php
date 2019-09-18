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

/**
 * Class Singleton
 */
final class Singleton
{

    /**
     *
     * @var self
     */
     private static $instance;

    /**
     *
     * @var string
     */
     public $name;

    /**
     *
     * @return Singleton
     */
    public static function getInstance() : self {
        if (!static::$instance instanceof self) {
            static::$instance = new self;
        }
        return static::$instance;
    }

    /**
     *
     * @param string $name
     */
    public function setName(string $name) {
        $this->name = $name;
    }

    /**
     * Singleton constructor.
     */
    private function __construct()
    {
        // TODO: constructor
    }

    private function __clone()
    {
        // TODO: Implement __clone() method.
    }

}

$singleton = Singleton::getInstance();
$test = Singleton::getInstance();
$singleton->setName("zhangsan");
$test->setName("lisi");

echo '$singleton : ' . $singleton->name .PHP_EOL;
echo '$test : ' . $test->name .PHP_EOL;
/**
 * result
 *
 * $singleton : lisi
 * $test : lisi
 */
