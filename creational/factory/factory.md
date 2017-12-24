# 工厂模式

工厂模式中，你不需要关系，每一种对象的具体创建过程，你只需指定你需要的类型，它内部会执行相应的函数返回所需的对象。例如，下面只需告诉它你要什么饮料，它就会返回相应的对象实例。注意：这些对象可能是完全不同的结构，包含完全不同的属性有不同的方法，但的实现同一接口。

``` go
func NewDrink(t Type) Interface {
	switch t {
	case juice:
		return Juice{...}
	case milk:
        return MakeMilk(...)
	default:
		return Water{...}
	}
}

func MakeMilk(...) Milk {
	...
    return Milk{}
}
```

**好处**： 使用代码的人不用关心及实现声明一个所需实例的过程，只需传入相应的参数即可得到所需的对象实例。

### 工厂模式 与 建造者模式

工厂模式和建造者模式都是提供参数得到对象，但是两者是完全不同的。在工厂模式中，当你给定参数，它就会创建一个完整的对象，而且给定相同的参数，得到的对象相同。而在建造者模式中，你给定的参数只创建了一个部件，给定的参数不完全相同会得到不同的对象，所以可以使用建造者模式定制不同的需要的对象。

### 二者结合

将 Car 改成 car 包

```go
func NewCar(t Type) Car {
	switch t {
	case bgb:
        return car.Car{}.Brand(builder.BMW).Color(builder.Green).Size(builder.Big).Build()
	case mrb:
        return car.Car{}.Brand(builder.MB).Color(builder.Red).Size(builder.Small).Build()
	default:
		return Car{}
	}
}
```
