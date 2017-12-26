# Functional Options

​	功能选项就是提供了多种功能可以选择。

​	在示例中

``` go
type Option func(*Options)
```

​	它将功能选项定义为函数

``` go
	for _, setter := range setters {
		setter(args)
    }
```

​	然后遍历选择的功能，执行这些函数。
