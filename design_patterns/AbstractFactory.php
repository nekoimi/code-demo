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

class Birds {
    public function do() { echo get_called_class() . PHP_EOL; }
}
class Chicken extends Birds {}
class Duck extends Birds {}

/**
 * Class AbstractFactory
 * 抽象工厂
 */
abstract class AbstractFactory
{
    abstract public function createBird ();
}

// 鸡类工厂
class ChickenFactory extends AbstractFactory {
    public function createBird()
    {
        return new Chicken();
    }
}
// 鸭类工程
class DuckFactory extends AbstractFactory {
    public function createBird()
    {
        return new Duck();
    }
}

$chickenFactory = new ChickenFactory();
$chicken = $chickenFactory->createBird();

$duckFactory = new DuckFactory();
$duck = $duckFactory->createBird();

$chicken->do();
$duck->do();
/**
 * result:
 *
 * Chicken
 * Duck
 */
