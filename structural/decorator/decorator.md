# 修饰者模式

修饰者模式，我们借助代码来理解。

``` go
type Object func(int) int
```

我们先定义了一个 Object 这样的一个类型。

``` go
func Decorate(fsource Object, fextend Object) Object {
	return func(n int) int {
		result := fextend(fsource(n))

		return result
	}
}
```

然后，写了一个修饰的规则。fextend(fsource(n))，也就是先执行 fsource()，将结果作为 fextend() 的参数。看具体例子。

``` go
func AddOne(n int) int {
	return n + 1
}

func Double(n int) int {
	return n * 2
}
```

这里定义了两个方法，一个是加一，一个是乘二。如果现在我又需要一个先加一再乘二的函数，就可以用 f1 := Decorate(AddOne, Double) 得到了。同理，需要加一再乘二最后还加一就可以用 f2 := Decorate(f1, AddOne) 这样的方法得到。

##### 注意：修饰者模式 与 策略模式 的对比，详见 [startegy.md](../../behavioral/strategy/strategy.md)
