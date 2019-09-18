### Golang的数组是值类型

- Slice -> 切片是数组的view

### Golang 切片

```go
arr := [...]int{0,1,2,3,4,5,6,7,8,9}
s1 := arr[2:6]  // 2,3,4,5
s2 := s1[3:5]   // 5,6
```

![](https://i.loli.net/2019/05/29/5cee7cd34dac721943.png)
