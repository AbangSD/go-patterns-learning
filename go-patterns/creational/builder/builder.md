# Builder

``` go
type Builder interface {
	Function1(parm) Builder
	Function2(parm) Builder
	Function3(parm) Builder
	Build() Interface
}
```

​	实现一个 Builder 接口，Builder 接口中包含若干个返回值是一个 Builder 接口的方法以及一个 Build 方法转换成一个 Interface 接口。

``` go
type Interface interface {
	Action1() error
	Action2() error
}
```

​	Interface 接口中包含若干个行为方法。
