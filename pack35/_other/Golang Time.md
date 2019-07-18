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
