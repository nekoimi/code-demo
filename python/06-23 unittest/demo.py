#!/usr/bin/python3
# -*- coding: utf-8 -*-

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
