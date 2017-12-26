# Bounded Parallelism Pattern

有界并行就是划分很多个执行的范围，然后让他们并发执行。

## 示例分析

​	**bounded_parallelism.go**: 来自https://github.com/tmrts/go-patterns/blob/master/concurrency/bounded_parallelism.go

```sh
	go run bounded_parallelism.go .
```

​	主要看 MD5All 中调用的 walkFiles 和 digester。

​	先看 walkFiles 函数，walkFiles 中主要用了 filepath.Walk 函数，下面这段代码有利于理解 filepath.Walk 的用法。

``` go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error{
		if err != nil {
			return err
		}
		fmt.Println("path:", path, "FileInfo.Name:", info.Name())
		return nil
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "walk failed with error: %v\n", err)
	}
}
```

``` go
func Walk(root string, walkFn WalkFunc) error

type WalkFunc func(path string, info os.FileInfo, err error) error
```

​	Walk 遍历在 root 中的所有文件，并且包括 root 在内的每一个文件、文件夹都会调用 walkFn。所有的访问文件、文件夹错误都通过 walkFn 筛选处理。walkFn 是 WalkFunc 类型的参数。

​	walkFiles 将所有文件的路径放到通道 paths 中，分好边界后，就开始并发执行。

```go
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}
```

​	digester 就只是简单的读取文件，读取文件然后 md5 加密放到一个 result 结构中，最后将结果放到通道 result 中。