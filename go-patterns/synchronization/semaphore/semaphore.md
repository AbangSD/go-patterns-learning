# Semaphore

semaphore/semaphore.go

``` go
type Interface interface {
	Acquire() error
	Release() error
}
```

``` go
type implementation struct {
	sem     chan struct{}
	timeout time.Duration
}
```

​	implementation 主要包含 sem 和 timeout。implementation 实现了 Interface 的接口。sem 是一个 chan 类型，并且带 n 个缓存，每次调用 Acquire() 都会往 sem 中放入一个 struct{}{}，当 sem 放满就会阻塞，如果阻塞至超时就会返回一个 error。Release() 会从 sem 中拿出一个 struct{}{}，如果 sem 中没有东西并等待至超时也会返回一个 error。

