# Fan-In && Fan-Out Messaging Pattern

​	Fan-In Messaging 和 Fan-Out Messaging 这两种模式正好实现了两种相反的功能。

​	Fan-In Messaging 实现了将多个相同的类型的通道中的数据汇总到一个通道中。而 Fan-Out Messaging 实现了将一个通道中的数据分发到多个通道中。

### Fan-In Messaging

​	这是一个 Fan-In Messaging 示例

``` go
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

    // 等发送结束后，关闭 out 通道，让它只能拿出数据不能放入数据
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
```

### Fan-Out Messaging

​	这是一个 Fan-Out Messaging 示例

``` go
func Split(ch <-chan int, n int) []<-chan int {
	cs := make([]chan int)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int))
	}

	distributeToChannels := func(ch <-chan int, cs []chan<- int) {
		// 分发完，把每个通道都关闭，让它们只能拿出数据不能放入数据
		defer func(cs []chan<- int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}
```