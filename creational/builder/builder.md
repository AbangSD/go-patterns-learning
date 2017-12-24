# 建造者模式

什么是建造者模式？

建造者模式又称为生成器模式，它是一种较为复杂、使用频率也相对较低的创建型模式。建造者模式为客户端返回的不是一个简单的产品，而是一个由多个部件组成的复杂产品，它将客户端与包含多个部件的复杂对象的创建过程分离，客户端无须知道复杂对象的内部组成部件与装配方式，只需要知道所需部件的类型即可。不同的具体部件定义了不同的创建过程，且具体的部件相互独立，增加新的部件非常方便，无须修改已有代码，系统具有较好的扩展性。

建造者模式有什么用，怎样写？这里我举了一个 car 的例子。

Builder 是一个这样的接口

``` go
type Builder interface {
	Brand(Brand) *Builder
	Color(Color) *Builder
	Size(Size) *Builder
	Build() Product
}
```

具体的 Car 是一个这样的结构

``` go
type Car struct {
	brand builder.Brand
	color builder.Color
	size  builder.Size
}
```

Car 分别实现了 Brand、Color、Size 和 Build 方法。Brand、Color、Size 方法的具体实现客户端无需知道，客户端只需要知道如何调用它，就可以生成不同想要的对象实例。