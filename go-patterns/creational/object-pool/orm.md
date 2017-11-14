# orm -- cockroach connection pool

**orm/cockroach/pool.go**



``` go
type Pool struct {
	lock sync.Mutex
	pool *ring.Ring
	size int // 实际长度
}
```

​	Pool 中的 pool 才是真正用于存放的 pool。它是一个 *container/ring.Ring 类型。它添加了一个 lock sync.Mutex 来保证操作的安全，size 是实际可用的连接数。



**func NewPool(db string, size int) *Pool**

​	先做一次判断保证 size 不超过 poolMaxSize，然后 new 一个 Pool。**pool.pool = ring.New(1)** 初始化 pool，否则报错**panic: runtime error: invalid memory address or nil pointer dereference** ，因为这里的 pool 是一个指针类型，不初始化值为 nil。最后，用一个 for 循环 size 次，每次用 **gorm.Open()** 获取一个数据库连接然后放到 pool 中。



**func (this *Pool) Close()**

​	让 Pool.pool 中的每个连接都关闭就可用关闭了。所以这里做了一个闭包函数，然后用 Ring.Do() 让每个部分都调用闭包函数。



**func (this *Pool) GetConnection() (orm.Connection, error)**

​	先判断是否有可用的连接 (size 是否为 0)，**conn := this.pool.Unlink(1)** 获取一个连接，最后 size 减 1.



**func (this *Pool) ReleaseConnection(v orm.Connection)**

​	**this.pool.Prev().Link(conn)** 将不用的连接放回 pool 中。



### 感谢

TechCatsLab: **https://github.com/TechCatsLab/Winston/tree/master/orm**