#!/usr/bin/python3
# -*- coding: utf-8 -*-
# ------------Oooo---
# -----------(----)---
# ------------)--/----
# ------------(_/-
# ----oooO----
# ----(---)----
# -----\--(--
# ------\_)-
# ----              千里之行， 始于足下.
#     author : Yprisoner <yyprisoner@gmail.com>
#     time : 2019/6/23 下午8:36
#                           ------
#    「 涙の雨が頬をたたくたびに美しく 」

import unittest


class TestMain(unittest.TestCase):

    def setUp(self) -> None:
        """
        每次方法之前执行
        """
        print("--- setUp ---")

    def tearDown(self) -> None:
        """
        每次方法之后执行
        """
        print("--- tearDown ---")

    def testMethod01(self):
        print('--- 测试方法一 Start ---')
        self.assertTrue(True)
        print('--- 测试方法一 End ---')

    def testMethod02(self):
        print('--- 测试方法二 Start ---')
        self.assertEqual(0, 0)
        print('--- 测试方法二 End ---')

    @classmethod
    def setUpClass(cls) -> None:
        print("--- Class Method setUpClass ---")

    @classmethod
    def tearDownClass(cls) -> None:
        print("--- Class Method tearDownClass ---")


if __name__ == '__main__':
    unittest.main()
