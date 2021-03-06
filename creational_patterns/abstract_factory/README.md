# Abstract Factory pattern

> 创建型模式 之 抽象工厂模式

抽象工厂模式（Abstract Factory Pattern）是一种比较常用的模式，

其定义如下：

Provide an interface for creating families of related or dependent objects without specifying
their concrete classes.

（为创建一组相关或相互依赖的对象提供一个接口，而且无须指定它们
的具体类。）


## 工厂方法模式的优点

- 首先，良好的封装性，代码结构清晰。

- 工厂方法模式的扩展性非常优秀。

- 屏蔽产品类。这一特点非常重要，产品类的实现如何变化，调用者都不需要关
  心，它只需要关心产品的接口，只要接口保持不变，系统中的上层模块就不要发生变化。
  
- 工厂方法模式是典型的解耦框架。高层模块值需要知道产品的抽象类，其他的实
  现类都不用关心