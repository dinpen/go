# github.com/srcio/go

本仓库中封装了一些 go 语言常用的工具函数，可以在日常开发中使用。

下面是一些包的使用示例。

## conc

- 并发运行多个函数，会开启相应数量的 goroutine

```go
conc.Run(func() {
    fmt.Println("Hello")
}, func() {
    fmt.Println("World")
})
```

- 并发运行多个函数，并且等待所有函数完成

```go
conc.Wait(func() {
    time.Sleep(1 * time.Second)
    fmt.Println("Hello")
}, func() {
    time.Sleep(4 * time.Second)
    fmt.Println("World")
})
```

- 并发运行多个函数，并且捕获 panic

```go
if p := conc.WaitAndRecover(func() {
fmt.Println("Hello")
panic("panic when print hello")
}, func() {
    fmt.Println("World")
    time.Sleep(1 * time.Second)
    panic("panic when print world")
}); p != nil {
    fmt.Printf("p.PanicInfo(): %v\n", p.PanicInfo())
    fmt.Printf("p.PanicVal(): %v\n", p.PanicVal())
}
```
