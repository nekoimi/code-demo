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

interface PrototypeInterface {
    public function copy();
}

class Prototype implements PrototypeInterface
{

    public function copy()
    {
         return clone $this;
    }
}