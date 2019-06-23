# UnitTest

再开发中， 单元测试必不可少.

在python中，使用python自带的unittest库进行测试

### 基本使用

使用实例:

```python
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
        print('--- 测试方法一 ---')

    def testMethod02(self):
        print('--- 测试方法二 ---')


if __name__ == '__main__':
    unittest.main()
```

输出结果:

```html
--- setUp ---
--- 测试方法一 ---
--- tearDown ---
--- setUp ---
--- 测试方法二 ---
--- tearDown ---
```

### 断言

可以看下 unittest.TestCase的源码

里面有很清晰的方法定义
