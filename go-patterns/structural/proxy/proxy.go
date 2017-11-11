package main

import (
	"fmt"
)

// 为了使用代理模式，必须让它们实现相同的方法
// 此处的接口只是为让我们更直接实现代理模式，只要他们都实现了此接口也就代表他们实现了代理模式
type IObject interface {
	Action(action string)
}

// 真正的对象
type Object struct{}

// Object 的方法
func (obj *Object) Action(action string) {
	fmt.Printf("%s", action)
}

// 让 ProxyObject 包含 1 个 *Object 类型的对象，让 ProxyObject 也可以调用 Object 的方法
type ProxyObject struct {
	object *Object
}

// 通过 ProxyObject.object 调用 Object 的方法，让 ProxyObject 实现与 Object 相同的方法
func (p *ProxyObject) Action(action string) {
	if p.object == nil {
		p.object = new(Object)
	}

	p.object.Action(action)
}

func main() {
	var po ProxyObject

	// 表面上是调用的是 ProxyObject 的 Action 方法，但事实上是调用的是 Object 的 Action 方法
	po.Action("hello world!")
}
