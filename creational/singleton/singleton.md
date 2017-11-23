# Singleton Pattern

``` go
type singleton struct {
	// ...
}

var (
	once sync.Once

	instance singleton
)

func New() singleton {
	once.Do(func() {
		instance = singleton{}
	})

	return instance
}
```

​	使用 sync.Once 中的 Do 来声明并初始化 instance，且 singleton 包外不可见，所以保证了只可能存在一个唯一的 singleton 类型的实例。

​	上面的解释看上去的确非常完美，但是事实上是存在问题的，让我们一起运行一下示例 singleton1，可以看到输出结果。

``` shell
s1 is {1} s2 is {0}
s1 is {1} s2 is {2}
```

​	从输出上看，once.Do 并没有保证初始化函数只执行一次。最终，产生了两个 singleton 实例。那正确的做法是怎么样的呢？

``` go
type singleton struct {
	// ...
}

var (
	once sync.Once

	instance *singleton
)

func New() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
```

​	再一起运行一下示例 singleton2，可以看到输出结果如下。

``` shell
s1 is &{1} s2 is &{1}
s1 is &{2} s2 is &{2}
&s1.N == &s2.N ? true
```

​	s1.N 和 s2.N 的地址相同表明 s1 和 s2 都指向了 instance。此时，实现了单例模式。

​	singleton3 是 map 的单例模式示例。