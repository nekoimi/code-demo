# Golang 的 time package

> go语言的时间格式化是真的骚

通常在php里面格式化时间直接使用php提供的date函数

```php
date('Y-m-d H:i:s')
```

格式化时间都是输入格式化参数， 很方便。

```go
time.Now().Format("2006-01-02 15:04:05")
```

时间格式化参数不是其他语言里边的 YYYY/MM/dd...这些

在format.go中写死的

就是这样写死的  搞得我一头雾水

```html
月份 1,01,Jan,January
日　 2,02,_2
时　 3,03,15,PM,pm,AM,am
分　 4,04
秒　 5,05
年　 06,2006
时区 -07,-0700,Z0700,Z07:00,-07:00,MST
周几 Mon,Monday
```
