# Factory

``` go
func NewObj(t Type) Obj {
	switch t {
	case Obj1:
		return Obj{...}
	case Obj2:
		return Obj{...}
	default:
		return Obj{...}
	}
}
```

​	**好处**： 使用代码的人不用关心及实现声明一个所需实例的过程，只需传入相应的参数即可得到所需的实例。

​	此处的 Obj 也可以是一个接口类型，具体示例参见 data。
