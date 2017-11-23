# Strategy Pattern

​	策略模式只是定义了一个宏观的规则，内部具体算法可能各不相同。

注意：

​	要保证 Operator 接口中的方法与 Operation 带有的方法一致。

## 与 Decorator 的区别

​	借用《head first 设计模式》中卖咖啡的修饰模式例子与维基百科中收税的策略模式例子对比一下。

​	修饰模式不会改变内部结构，只是在原有基础上扩充。不同的修饰后他们本质上是同一类东西，但是各自的属性和方法不太一样，让他们看起来是不同的。而策略模式只是定义了一个外部框架，内部属性方法不同。内部不同的类型在策略这种模式下看起来是同一种类型。

​	例如，在卖咖啡的例子中，加牛奶的咖啡，加糖的咖啡，加牛奶和糖的咖啡本质上都是咖啡，所以定义了一个普通的咖啡方法得到普通咖啡，然后分别定义了加牛奶加糖两个方法，通过给普通咖啡修饰就可以得到不同种类的咖啡。虽然这几种咖啡实质都是普通咖啡，但是他们名字和价格都不一样，看起来是不同的。

​	而在收税的例子中，中国人美国人的税在需要纳的数目种类比例都不相同，但最终他们都是调用 Taxes()，虽然内部完全不同，但是最终形式结果保持同样，看起来是相同的。

​	代码详见 **./compare/coffee** 与 **./compare/taxes** 